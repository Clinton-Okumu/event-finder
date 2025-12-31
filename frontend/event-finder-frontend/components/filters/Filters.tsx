"use client";

import { useState } from "react";
import CategoryFilter from "./CategoryFilter";
import DateFilter from "./DateFilter";
import PriceRangeFilter from "./PriceRangeFilter";
import SearchBar from "./SearchBar";
import { Button } from "@/components/ui/button";
import { X } from "lucide-react";

export interface FilterValues {
  search: string;
  categories: string[];
  date: "today" | "weekend" | "month" | "custom" | null;
  customDate: string;
  maxPrice: number;
}

interface FiltersProps {
  onFiltersChange: (filters: FilterValues) => void;
  filters?: FilterValues;
}

export default function Filters({ onFiltersChange, filters }: FiltersProps) {
    const [search, setSearch] = useState(filters?.search || "");
    const [categories, setCategories] = useState<string[]>(filters?.categories || []);
    const [date, setDate] = useState<"today" | "weekend" | "month" | "custom" | null>(filters?.date || null);
    const [customDate, setCustomDate] = useState(filters?.customDate || "");
    const [maxPrice, setMaxPrice] = useState(filters?.maxPrice || 2000);

    const handleApplyFilters = () => {
        onFiltersChange({
            search,
            categories,
            date,
            customDate,
            maxPrice,
        });
    };

    const handleClearFilters = () => {
        setSearch("");
        setCategories([]);
        setDate(null);
        setCustomDate("");
        setMaxPrice(2000);
        onFiltersChange({
            search: "",
            categories: [],
            date: null,
            customDate: "",
            maxPrice: 2000,
        });
    };

    return (
        <aside className="w-72 bg-card border border-border rounded-xl p-6 shadow-lg h-fit sticky top-6">
            <div className="flex items-center justify-between mb-4">
                <h2 className="text-xl font-semibold text-foreground">Filters</h2>
                <Button
                    variant="ghost"
                    size="icon"
                    className="h-6 w-6"
                    onClick={handleClearFilters}
                >
                    <X className="h-4 w-4" />
                </Button>
            </div>
            <p className="text-sm text-muted-foreground mb-6">Refine your search</p>

            <SearchBar value={search} onChange={setSearch} />

            <DateFilter
                value={date}
                customDate={customDate}
                onChange={setDate}
                onCustomDateChange={setCustomDate}
            />

            <CategoryFilter
                selectedCategories={categories}
                onChange={setCategories}
            />

            <PriceRangeFilter value={maxPrice} onChange={setMaxPrice} />

            <Button className="w-full mt-6" size="lg" onClick={handleApplyFilters}>
                Apply Filters
            </Button>
        </aside>
    );
}
