from ..extensions import db
from sqlalchemy import Numeric


class BookingItem(db.Model):
    __tablename__ = "booking_items"
    id = db.Column(db.Integer, primary_key=True)
    booking_id = db.Column(db.Integer, db.ForeignKey("bookings.id"), nullable=False)
    booking = db.relationship("Booking", backref="items")
    ticket_type_id = db.Column(
        db.Integer, db.ForeignKey("event_tickets.id"), nullable=False
    )
    ticket_type = db.relationship("EventTicket")
    quantity = db.Column(db.Integer, nullable=False)
    price_at_purchase = db.Column(Numeric(10, 2), nullable=False)


@property
def total_price(self):
    return self.quantity * self.price_at_purchase
