"use client";

interface PriceRangeFilterProps {
  value: number;
  onChange: (value: number) => void;
}

export default function PriceRangeFilter({ value, onChange }: PriceRangeFilterProps) {
    return (
        <div className="mb-6">
            <h3 className="font-semibold mb-3 text-foreground">Price Range</h3>

            <input
                type="range"
                min={0}
                max={2000}
                value={value}
                onChange={(e) => onChange(Number(e.target.value))}
                className="w-full accent-primary"
            />

            <div className="flex justify-between text-sm mt-2 text-muted-foreground">
                <span>ksh.0</span>
                <span>ksh.{value}</span>
                <span>ksh.2000+</span>
            </div>
        </div>
    );
}
