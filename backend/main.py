import os
import click
from app import create_app
from app.extensions import db
from app.models.user import User

app = create_app()


@app.cli.command("create-admin")
@click.argument("email")
@click.argument("password")
def create_admin(email, password):
    """Create an admin user."""
    with app.app_context():
        existing = User.query.filter_by(email=email).first()
        if existing:
            click.echo(f"User with email {email} already exists.")
            return

        user = User(email=email, role="admin")
        user.set_password(password)
        db.session.add(user)
        db.session.commit()
        click.echo(f"Admin user {email} created successfully.")


if __name__ == "__main__":
    port = int(os.getenv("FLASK_RUN_PORT", 5000))
    debug = os.getenv("FLASK_DEBUG", "true").lower() == "true"

    app.run(debug=debug, port=port)
