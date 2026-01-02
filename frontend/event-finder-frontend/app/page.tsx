"use client";

import { useState, useEffect } from "react";
import Events from "@/components/events/Events";
import Filters, { FilterValues } from "@/components/filters/Filters";
import PageHeader from "@/components/layout/PageHeader";
import { Button } from "@/components/ui/button";
import { ArrowRight } from "lucide-react";
import Link from "next/link";
import { Event } from "@/lib/types/events";
import { getEvents } from "@/lib/api/events";

export default function Home() {
  const [allEvents, setAllEvents] = useState<Event[]>([]);
  const [filteredEvents, setFilteredEvents] = useState<Event[]>([]);
  const [loading, setLoading] = useState(true);
  const [filters, setFilters] = useState<FilterValues>({
    search: "",
    categories: [],
    date: null,
    customDate: "",
    maxPrice: 2000,
  });

  useEffect(() => {
    fetchEvents();
  }, []);

  useEffect(() => {
    applyFilters();
  }, [filters, allEvents]);

  const fetchEvents = async () => {
    try {
      setLoading(true);
      const data = await getEvents();
      setAllEvents(data);
    } catch (error) {
      console.error("Failed to fetch events:", error);
    } finally {
      setLoading(false);
    }
  };

  const applyFilters = () => {
    let filtered = [...allEvents];

    if (filters.search) {
      filtered = filtered.filter(
        (event) =>
          event.title.toLowerCase().includes(filters.search.toLowerCase()) ||
          event.description.toLowerCase().includes(filters.search.toLowerCase()) ||
          event.location.toLowerCase().includes(filters.search.toLowerCase())
      );
    }

    if (filters.categories.length > 0) {
      filtered = filtered.filter((event) => {
        const eventCategory = event.description.toLowerCase();
        return filters.categories.some((cat) =>
          eventCategory.includes(cat.toLowerCase())
        );
      });
    }

    if (filters.date || filters.customDate) {
      const today = new Date();
      today.setHours(0, 0, 0, 0);

      filtered = filtered.filter((event) => {
        const evtDate = new Date(event.start_time);

        if (filters.customDate) {
          const filterDate = new Date(filters.customDate);
          filterDate.setHours(0, 0, 0, 0);
          return evtDate.getTime() === filterDate.getTime();
        }

        if (filters.date === "today") {
          return evtDate.getTime() === today.getTime();
        }

        if (filters.date === "weekend") {
          const dayOfWeek = evtDate.getDay();
          return dayOfWeek === 0 || dayOfWeek === 6;
        }

        if (filters.date === "month") {
          const nextMonth = new Date(today);
          nextMonth.setMonth(nextMonth.getMonth() + 1);
          return evtDate >= today && evtDate < nextMonth;
        }

        return true;
      });
    }

    if (filters.maxPrice < 2000) {
      filtered = filtered.filter(
        (event) =>
          event.price === undefined || event.price <= filters.maxPrice
      );
    }

    setFilteredEvents(filtered);
  };

  const latestEvents = filteredEvents.slice(0, 6);

  if (loading) {
    return (
      <div className="container mx-auto px-4 py-8">
        <div className="text-center py-12">
          <p className="text-muted-foreground">Loading events...</p>
        </div>
      </div>
    );
  }

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex flex-col lg:flex-row gap-8">
        <div className="lg:w-72 shrink-0">
          <Filters onFiltersChange={setFilters} filters={filters} />
        </div>

        <main className="flex-1 min-w-0">
          <PageHeader
            title="Discover Amazing Events"
            subtitle={`Found ${filteredEvents.length} events`}
          />

          <Events events={latestEvents} />
          <div className="flex mt-10 justify-center items-center">
            <Link href="/events">
              <Button>
                More Events
                <ArrowRight className="ml-1 h-4 w-4" />
              </Button>
            </Link>
          </div>
        </main>
      </div>
    </div>
  );
}
