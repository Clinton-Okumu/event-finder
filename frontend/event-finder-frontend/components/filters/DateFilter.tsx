import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";

export default function DateFilter() {
    return (
        <div className="mb-6">
            <h3 className="font-semibold mb-3 text-foreground">Date</h3>

            <div className="flex gap-2 mb-3">
                <Button variant="default" size="sm">Today</Button>
                <Button variant="outline" size="sm">Weekend</Button>
            </div>

            <Button variant="outline" size="sm" className="w-full mb-3">
                This Month
            </Button>

            <Input type="date" />
        </div>
    );
}
