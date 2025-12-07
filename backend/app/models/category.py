from ..extensions import db


class Category(db.Model):
    __tablename__ = "category"
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(255), nullable=False)
    description = db.Column(db.String(255), nullable=False)
    slug = db.Column(db.String(255), nullable=False)
    events_id = db.Column(db.Integer, db.ForeignKey("events.id"))
    events = db.relationship("Events", backref="categories")
