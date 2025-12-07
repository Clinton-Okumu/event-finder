from flask import Blueprint, request, jsonify
from ..extensions import db
from ..models.events import Event
from ..utils.auth import admin_required
from flask_jwt_extended import get_jwt_identity
from datetime import datetime

events = Blueprint("events", __name__, url_prefix="/events")


@events.route("/", methods=["GET"])
def get_events():
    all_events = Event.query.all()
    return jsonify([e.to_dict() for e in all_events]), 200


@events.route("/<int:event_id>", methods=["GET"])
def get_event(event_id: int):
    event = Event.query.get_or_404(event_id)
    return jsonify(event.to_dict()), 200


@events.route("/", methods=["POST"])
@admin_required
def create_event():
    data: dict = request.get_json() or {}

    start_str: str = data.get("start_time", "")
    end_str: str = data.get("end_time", "")
    start_time: datetime = (
        datetime.fromisoformat(start_str) if start_str else datetime.now()
    )
    end_time: datetime = datetime.fromisoformat(end_str) if end_str else datetime.now()

    event = Event(
        title=data.get("title", ""),
        description=data.get("description", ""),
        location=data.get("location", ""),
        start_time=start_time,
        end_time=end_time,
        category_id=data.get("category_id", 1),
        image_url=data.get("image_url", ""),
        total_tickets=data.get("total_tickets", 0),
        tickets_remaining=data.get("tickets_remaining", 0),
        created_by=get_jwt_identity(),
    )

    db.session.add(event)
    db.session.commit()

    return jsonify(event.to_dict()), 201


@events.route("/<int:event_id>", methods=["PUT"])
@admin_required
def update_event(event_id: int):
    event = Event.query.get_or_404(event_id)
    data: dict = request.get_json() or {}

    if "start_time" in data:
        event.start_time = datetime.fromisoformat(data["start_time"])
    if "end_time" in data:
        event.end_time = datetime.fromisoformat(data["end_time"])

    event.title = data.get("title", event.title)
    event.description = data.get("description", event.description)
    event.location = data.get("location", event.location)
    event.category_id = data.get("category_id", event.category_id)
    event.image_url = data.get("image_url", event.image_url)
    event.total_tickets = data.get("total_tickets", event.total_tickets)
    event.tickets_remaining = data.get("tickets_remaining", event.tickets_remaining)

    db.session.commit()
    return jsonify(event.to_dict()), 200


@events.route("/<int:event_id>", methods=["DELETE"])
@admin_required
def delete_event(event_id: int):
    event = Event.query.get_or_404(event_id)
    db.session.delete(event)
    db.session.commit()
    return jsonify({"message": "Event deleted"}), 200
