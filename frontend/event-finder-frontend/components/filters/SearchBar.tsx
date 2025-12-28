import { Search } from "lucide-react";
import { InputGroup, InputGroupAddon, InputGroupInput } from "@/components/ui/input-group";

export default function SearchBar() {
    return (
        <div className="mb-6">
            <InputGroup>
                <InputGroupAddon align="inline-start">
                    <Search className="h-4 w-4" />
                </InputGroupAddon>
                <InputGroupInput
                    type="text"
                    placeholder="Search events..."
                />
            </InputGroup>
        </div>
    );
}
