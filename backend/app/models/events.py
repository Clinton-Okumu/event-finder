from ..extensions import db
from sqlalchemy import Enum
from datetime import datetime
from sqlalchemy.orm import mapped_column, Mapped


class Event(db.Model):
    __tablename__ = "events"

    id: Mapped[int] = mapped_column(primary_key=True)
    title: Mapped[str] = mapped_column(nullable=False)
    description: Mapped[str] = mapped_column(db.Text, nullable=False)
    location: Mapped[str] = mapped_column(nullable=False)
    start_time: Mapped[datetime] = mapped_column(db.DateTime, nullable=False)
    end_time: Mapped[datetime] = mapped_column(db.DateTime, nullable=False)
    created_by: Mapped[int] = mapped_column(
        db.Integer, db.ForeignKey("users.id"), nullable=False
    )
    status: Mapped[str] = mapped_column(
        Enum("pending", "confirmed", "canceled", name="event_status"),
        default="pending",
        nullable=False,
    )
    image_url: Mapped[str] = mapped_column(nullable=False)
    total_tickets: Mapped[int] = mapped_column(nullable=False)
    tickets_remaining: Mapped[int] = mapped_column(nullable=False)
    category_id: Mapped[int] = mapped_column(
        db.Integer, db.ForeignKey("categories.id"), nullable=False
    )
    created_at: Mapped[datetime] = mapped_column(
        db.DateTime, default=datetime.now, nullable=False
    )
    updated_at: Mapped[datetime] = mapped_column(
        db.DateTime, default=datetime.now, onupdate=datetime.now, nullable=False
    )

    creator = db.relationship("User", backref="events")
    category = db.relationship("Category", backref="events")

    def to_dict(self) -> dict:
        return {
            "id": self.id,
            "title": self.to_dict(),
            "description": self.description,
            "location": self.location,
            "start_time": self.start_time.isoformat(),
            "end_time": self.end_time.isoformat(),
            "created_by": self.created_by,
            "status": self.status,
            "image_url": self.image_url,
            "total_tickets": self.total_tickets,
            "tickets_remaining": self.tickets_remaining,
            "category_id": self.category_id,
        }
