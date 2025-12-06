from flask import Blueprint, request, jsonify
from ..extensions import db
from ..models.user import User
from flask_jwt_extended import create_access_token, jwt_required, get_jwt_identity

auth = Blueprint("auth", __name__, url_prefix="/auth")


@auth.post("/register")
@jwt_required(optional=True)
def register():
    data = request.get_json() or {}
    email = data.get("email")
    password = data.get("password")
    role = data.get("role", "user")

    if not email or not password:
        return jsonify({"error": "Email and password required"}), 400

    # Prevent admin creation unless authenticated as admin
    if role == "admin":
        user_id = get_jwt_identity()
        if not user_id:
            return jsonify({"error": "Login as admin to create admin users"}), 403

        current_user = User.query.get(user_id)
        if not current_user or not current_user.is_admin:
            return jsonify({"error": "Only admins can create admin accounts"}), 403

    # Check if email exists
    if User.query.filter_by(email=email).first():
        return jsonify({"error": "Email already registered"}), 400
    # Create user
    user = User(email=email, role=role)
    user.password = password

    db.session.add(user)
    db.session.commit()

    return jsonify({"message": "User registered successfully"}), 201


@auth.post("/login")
def login():
    data = request.get_json() or {}
    email = data.get("email")
    password = data.get("password")

    user = User.query.filter_by(email=email).first()
    if not user or not user.check_password(password):
        return jsonify({"error": "Invalid credentials"}), 401

    token = create_access_token(identity=user.id)

    return (
        jsonify(
            {
                "access_token": token,
                "user": {"id": user.id, "email": user.email, "role": user.role},
            }
        ),
        200,
    )
