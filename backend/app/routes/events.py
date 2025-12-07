from flask import Blueprint, request, jsonify
from ..extensions import db
from ..models.events import Event
from ..utils.auth import admin_required

events = Blueprint("events", __name__, url_prefix="/events")


@events.route("/", methods=["GET"])
def get_events():
    events = Event.query.all()
    data = [event.to_dict() for event in events]
    return jsonify(data), 200


@events.route("/<int:event_id>", methods=["GET"])
def get_event(event_id):
    event = Event.query.get(event_id)
    if not event:
        return jsonify({"error": "Event not found"}), 404
    return jsonify(event.to_dict()), 200


@events.route("/", methods=["POST"])
@admin_required
def create_event():
    data = request.get_json()

    event = Event(
        title=data.get("title"),
        description=data.get("description"),
        location=data.get("location"),
        date=data.get("date"),
        category_id=data.get("category_id"),
    )
    db.session.add(event)
    db.session.commit()

    return jsonify(event.to_dict()), 201


@events.route("/<int:event_id>", methods=["PUT"])
@admin_required
def update_event(event_id):
    event = Event.query.get(event_id)
    if not event:
        return jsonify({"error": "Event not found"}), 404

    data = request.get_json()

    event.title = data.get("title", event.title)
    event.description = data.get("description", event.description)
    event.location = data.get("location", event.location)
    event.date = data.get("date", event.date)
    event.category_id = data.get("category_id", event.category_id)

    db.session.commit()
    return jsonify(event.to_dict()), 200


@events.route("/<int:event_id>", methods=["DELETE"])
@admin_required
def delete_event(event_id):
    event = Event.query.get(event_id)
    if not event:
        return jsonify({"error": "Event not found"}), 404

    db.session.delete(event)
    db.session.commit()
    return jsonify({"message": "Event deleted"}), 200
