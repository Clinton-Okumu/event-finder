export default function PriceRangeFilter() {
    return (
        <div className="mb-6">
            <h3 className="font-semibold mb-3 text-foreground">Price Range</h3>

            <input
                type="range"
                min={0}
                max={500}
                defaultValue={200}
                className="w-full accent-primary"
            />

            <div className="flex justify-between text-sm mt-2 text-muted-foreground">
                <span>ksh.0</span>
                <span>ksh.2000+</span>
            </div>
        </div>
    );
}
