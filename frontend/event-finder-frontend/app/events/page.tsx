"use client";

import EventCard from "@/components/events/EventCard";
import Filters, { FilterValues } from "@/components/filters/Filters";
import PageHeader from "@/components/layout/PageHeader";
import { Button } from "@/components/ui/button";
import { getEvents } from "@/lib/api/events";
import { Event } from "@/lib/types/events";
import { ChevronLeft, ChevronRight } from "lucide-react";
import { useEffect, useState } from "react";

const EVENTS_PER_PAGE = 6;

export default function AllEventsPage() {
  const [events, setEvents] = useState<Event[]>([]);
  const [loading, setLoading] = useState(true);
  const [currentPage, setCurrentPage] = useState(1);
  const [filteredEvents, setFilteredEvents] = useState<Event[]>([]);
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
  }, [filters, events]);

  const fetchEvents = async () => {
    try {
      setLoading(true);
      const data = await getEvents();
      setEvents(data);
    } catch (error) {
      console.error("Failed to fetch events:", error);
    } finally {
      setLoading(false);
    }
  };

  const applyFilters = () => {
    let filtered = [...events];

    if (filters.search) {
      filtered = filtered.filter(
        (event) =>
          event.title.toLowerCase().includes(filters.search.toLowerCase()) ||
          event.description
            .toLowerCase()
            .includes(filters.search.toLowerCase()) ||
          event.location.toLowerCase().includes(filters.search.toLowerCase()),
      );
    }

    if (filters.categories.length > 0) {
      filtered = filtered.filter((event) => {
        const eventCategory = event.description.toLowerCase();
        return filters.categories.some((cat) =>
          eventCategory.includes(cat.toLowerCase()),
        );
      });
    }

    if (filters.date || filters.customDate) {
      const today = new Date();
      today.setHours(0, 0, 0, 0);

      const eventDate = (dateStr: string) => {
        const [month, day, year] = dateStr.split("/").map(Number);
        return new Date(year, month - 1, day);
      };

      filtered = filtered.filter((event) => {
        const evtDate = eventDate(event.date);

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
        (event) => event.price === undefined || event.price <= filters.maxPrice,
      );
    }

    setFilteredEvents(filtered);
    setCurrentPage(1);
  };

  const totalPages = Math.ceil(filteredEvents.length / EVENTS_PER_PAGE);
  const startIndex = (currentPage - 1) * EVENTS_PER_PAGE;
  const endIndex = startIndex + EVENTS_PER_PAGE;
  const currentEvents = filteredEvents.slice(startIndex, endIndex);

  const handlePageChange = (page: number) => {
    setCurrentPage(page);
    window.scrollTo({ top: 0, behavior: "smooth" });
  };

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
            title="All Events"
            subtitle={`This are all the events that are available on Event-Finder`}
          />

          {filteredEvents.length === 0 ? (
            <div className="text-center py-12">
              <p className="text-muted-foreground text-lg">No events found</p>
            </div>
          ) : (
            <>
              <div className="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-6">
                {currentEvents.map((event) => (
                  <EventCard key={event.id} {...event} />
                ))}
              </div>

              {totalPages > 1 && (
                <div className="flex justify-center items-center gap-2 mt-10">
                  <Button
                    variant="outline"
                    size="icon"
                    onClick={() => handlePageChange(currentPage - 1)}
                    disabled={currentPage === 1}
                  >
                    <ChevronLeft className="h-4 w-4" />
                  </Button>

                  <div className="flex items-center gap-2">
                    {Array.from({ length: totalPages }, (_, i) => i + 1).map(
                      (page) => (
                        <Button
                          key={page}
                          variant={currentPage === page ? "default" : "outline"}
                          size="icon"
                          onClick={() => handlePageChange(page)}
                        >
                          {page}
                        </Button>
                      ),
                    )}
                  </div>

                  <Button
                    variant="outline"
                    size="icon"
                    onClick={() => handlePageChange(currentPage + 1)}
                    disabled={currentPage === totalPages}
                  >
                    <ChevronRight className="h-4 w-4" />
                  </Button>
                </div>
              )}

              <div className="text-center mt-4 text-sm text-muted-foreground">
                Page {currentPage} of {totalPages}
              </div>
            </>
          )}
        </main>
      </div>
    </div>
  );
}
