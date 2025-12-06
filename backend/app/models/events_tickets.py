from ..extensions import db
from sqlalchemy import Enum, Numeric


class EventTicket(db.Model):
    __tablename__ = "event_tickets"
    id = db.Column(db.Integer, primary_key=True)
    event_id = db.Column(db.Integer, db.ForeignKey("events.id"), nullable=False)
    event = db.relationship("Event", backref="tickets")
    name = db.Column(
        Enum("Early Bird", "VIP", "Regular", "VVIP", name="ticket_type"), nullable=False
    )
    price = db.Column(Numeric(10, 2), nullable=False)
    quantity = db.Column(db.Integer, nullable=False)
    description = db.Column(db.Text, nullable=True)
