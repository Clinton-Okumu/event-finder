"use client";

import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";

type DateFilterValue = "today" | "weekend" | "month" | "custom";

interface DateFilterProps {
  value: DateFilterValue | null;
  customDate: string;
  onChange: (value: DateFilterValue | null) => void;
  onCustomDateChange: (date: string) => void;
}

export default function DateFilter({ value, customDate, onChange, onCustomDateChange }: DateFilterProps) {
    return (
        <div className="mb-6">
            <h3 className="font-semibold mb-3 text-foreground">Date</h3>

            <div className="flex gap-2 mb-3">
                <Button
                    variant={value === "today" ? "default" : "outline"}
                    size="sm"
                    onClick={() => onChange(value === "today" ? null : "today")}
                >
                    Today
                </Button>
                <Button
                    variant={value === "weekend" ? "default" : "outline"}
                    size="sm"
                    onClick={() => onChange(value === "weekend" ? null : "weekend")}
                >
                    Weekend
                </Button>
            </div>

            <Button
                variant={value === "month" ? "default" : "outline"}
                size="sm"
                className="w-full mb-3"
                onClick={() => onChange(value === "month" ? null : "month")}
            >
                This Month
            </Button>

            <Input
                type="date"
                value={customDate}
                onChange={(e) => {
                    onCustomDateChange(e.target.value);
                    onChange(e.target.value ? "custom" : null);
                }}
            />
        </div>
    );
}
