from flask import Blueprint, request, jsonify
from ..extensions import db
from ..models.user import User
from flask_jwt_extended import create_access_token

auth = Blueprint("auth", __name__, url_prefix="/auth")


@auth.post("/register")
def register():
    data = request.get_json()
    email = data.get("email")
    password = data.get("password")

    if not email or not password:
        return jsonify({"error": "Missing email or password"}), 400

    existing = User.query.filter_by(email=email).first()
    if existing:
        return jsonify({"error": "Email already registered"}), 400

    # Create the User
    user = User(email=email)
    user.set_password(password)

    db.session.add(user)
    db.session.commit()

    return jsonify({"message": "Successfully registered"}), 201


@auth.post("/login")
def login():
    data = request.get_json()
    email = data.get("email")
    password = data.get("password")

    user = User.query.filter_by(email=email).first()
    if not user or not user.check_password(password):
        return jsonify({"error": "Invalid credentials"}), 401

    token = create_access_token(identity=user.id)

    return (
        jsonify({"access_token": token, "user": {"id": user.id, "email": user.email}}),
        200,
    )
