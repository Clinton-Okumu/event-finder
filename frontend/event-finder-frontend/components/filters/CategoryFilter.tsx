import { Label } from "@/components/ui/label";

const categories = ["Music", "Sports", "Theater", "Comedy"];

export default function CategoryFilter() {
    return (
        <div className="mb-6">
            <h3 className="font-semibold mb-3 text-foreground">Category</h3>

            <div className="flex flex-col gap-2">
                {categories.map((cat) => (
                    <Label key={cat} className="flex items-center gap-2 cursor-pointer">
                        <input type="checkbox" className="accent-primary" />
                        <span className="text-sm">{cat}</span>
                    </Label>
                ))}
            </div>
        </div>
    );
}
