"use client";

import { Button } from "@/components/ui/button";
import { ArrowRight } from "lucide-react";
import Link from "next/link";

export default function AboutHero() {
  return (
    <section className="py-10 text-center">
      <div className="max-w-3xl mx-auto space-y-6">
        <h1 className="text-5xl md:text-6xl font-bold tracking-tight">
          About <span className="text-primary">Event Finder</span>
        </h1>
        <p className="text-xl text-muted-foreground">
          Connecting people with unforgettable experiences since 2025
        </p>
        <div className="flex justify-center gap-4 pt-4">
          <Link href="/events">
            <Button size="lg">
              Explore Events
              <ArrowRight className="ml-2 h-5 w-5" />
            </Button>
          </Link>
        </div>
      </div>
    </section>
  );
}
