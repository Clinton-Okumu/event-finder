import Events from "@/components/events/Events";
import Filters from "@/components/filters/Filters";
import PageHeader from "@/components/layout/PageHeader";
import { Button } from "@/components/ui/button";
import { getEvents } from "@/lib/api/events";
import { ArrowRight } from "lucide-react";

export default async function Home() {
  const events = await getEvents();

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex flex-col lg:flex-row gap-8">
        <div className="lg:w-72 shrink-0">
          <Filters />
        </div>

        <main className="flex-1 min-w-0">
          <PageHeader
            title="Discover Amazing Events"
            subtitle="Find your next experience"
          />

          <Events events={events} />
          <div className="flex mt-10 justify-center items-center">
            <Button>
              More Events
              <ArrowRight className="ml-1 h-4 w-4" />
            </Button>
          </div>
        </main>
      </div>
    </div>
  );
}
