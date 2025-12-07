from ..extensions import db, bcrypt
from sqlalchemy.orm import Mapped, mapped_column


class User(db.Model):
    __tablename__ = "users"

    id: Mapped[int] = mapped_column(primary_key=True)
    email: Mapped[str] = mapped_column(nullable=False, unique=True)
    _password: Mapped[str] = mapped_column(nullable=False)
    role: Mapped[str] = mapped_column(nullable=False, default="user")

    def __init__(self, email: str, password: str, role: str = "user"):
        self.email = email.lower()  # normalize email
        self.password = password  # setter hashes it
        self.role = role

    @property
    def password(self) -> None:
        raise AttributeError("Password is write-only")

    @password.setter
    def password(self, raw_password: str) -> None:
        self._password = bcrypt.generate_password_hash(raw_password).decode("utf-8")

    def check_password(self, raw_password: str) -> bool:
        return bcrypt.check_password_hash(self._password, raw_password)

    @property
    def is_admin(self) -> bool:
        return self.role == "admin"

    def to_dict(self) -> dict:
        return {"id": self.id, "email": self.email, "role": self.role}
