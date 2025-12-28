"use client";

import CategoryFilter from "./CategoryFilter";
import DateFilter from "./DateFilter";
import PriceRangeFilter from "./PriceRangeFilter";
import SearchBar from "./SearchBar";
import { Button } from "@/components/ui/button";

export default function Filters() {
    return (
        <aside className="w-72 bg-card border border-border rounded-xl p-6 shadow-sm h-fit sticky top-6">
            <h2 className="text-xl font-semibold mb-4 text-foreground">Filters</h2>
            <p className="text-sm text-muted-foreground mb-6">Refine your search</p>

            <SearchBar />

            <DateFilter />

            <CategoryFilter />

            <PriceRangeFilter />

            <Button className="w-full mt-6" size="lg">
                Apply Filters
            </Button>
        </aside>
    );
}
