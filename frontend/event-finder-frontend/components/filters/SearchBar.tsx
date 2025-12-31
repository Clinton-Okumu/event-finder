"use client";

import { Search } from "lucide-react";
import { InputGroup, InputGroupAddon, InputGroupInput } from "@/components/ui/input-group";

interface SearchBarProps {
  value: string;
  onChange: (value: string) => void;
}

export default function SearchBar({ value, onChange }: SearchBarProps) {
    return (
        <div className="mb-6">
            <InputGroup>
                <InputGroupAddon align="inline-start">
                    <Search className="h-4 w-4" />
                </InputGroupAddon>
                <InputGroupInput
                    type="text"
                    placeholder="Search events..."
                    value={value}
                    onChange={(e) => onChange(e.target.value)}
                />
            </InputGroup>
        </div>
    );
}
