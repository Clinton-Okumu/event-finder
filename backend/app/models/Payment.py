from ..extensions import db
from sqlalchemy import Enum, Numeric
from datetime import datetime


class Payment(db.Model):
    __tablename__ = "payments"

    id = db.Column(db.Integer, primary_key=True)
    booking_id = db.Column(db.Integer, db.ForeignKey("bookings.id"), nullable=False)
    booking = db.relationship("Booking", backref="payments")
    method = db.Column(db.String(50), nullable=False, default="card")
    amount = db.Column(Numeric(10, 2), nullable=False)
    status = db.Column(
        Enum("pending", "completed", "failed", "refunded", name="payment_status"),
        default="pending",
        nullable=False,
    )
    transaction_id = db.Column(db.String(255), nullable=True, unique=True)
    created_at = db.Column(db.DateTime, default=datetime.now(), nullable=False)
    updated_at = db.Column(
        db.DateTime, default=datetime.now(), onupdate=datetime.now(), nullable=False
    )
