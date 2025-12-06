from ..extensions import db
from sqlalchemy import Enum
from datetime import datetime


class Booking(db.Model):
    __tablename__ = "bookings"
    id = db.Column(db.Integer, primary_key=True)
    user_id = db.Column(db.Integer, db.ForeignKey("users.id"), nullable=False)
    user = db.relationship("User", backref="bookings")
    event_id = db.Column(db.Integer, db.ForeignKey("events.id"), nullable=False)
    event = db.relationship("Event", backref="bookings")
    booking_status = db.Column(
        Enum("pending", "confirmed", "canceled", name="booking_status"),
        default="pending",
        nullable=False,
    )
    created_at = db.Column(db.DateTime, default=datetime.now(), nullable=False)
    updated_at = db.Column(
        db.DateTime, default=datetime.now(), onupdate=datetime.now(), nullable=False
    )
