"use client";

import { Label } from "@/components/ui/label";

const categories = ["Music", "Sports", "Theater", "Comedy"];

interface CategoryFilterProps {
  selectedCategories: string[];
  onChange: (categories: string[]) => void;
}

export default function CategoryFilter({ selectedCategories, onChange }: CategoryFilterProps) {
    const handleCategoryChange = (category: string, checked: boolean) => {
        if (checked) {
            onChange([...selectedCategories, category]);
        } else {
            onChange(selectedCategories.filter((cat) => cat !== category));
        }
    };

    return (
        <div className="mb-6">
            <h3 className="font-semibold mb-3 text-foreground">Category</h3>

            <div className="flex flex-col gap-2">
                {categories.map((cat) => (
                    <Label key={cat} className="flex items-center gap-2 cursor-pointer">
                        <input
                            type="checkbox"
                            className="accent-primary"
                            checked={selectedCategories.includes(cat)}
                            onChange={(e) => handleCategoryChange(cat, e.target.checked)}
                        />
                        <span className="text-sm">{cat}</span>
                    </Label>
                ))}
            </div>
        </div>
    );
}
