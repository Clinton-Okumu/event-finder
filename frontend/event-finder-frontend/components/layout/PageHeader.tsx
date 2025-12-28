interface PageHeaderProps {
    title: string;
    subtitle?: string;
}

export default function PageHeader({ title, subtitle }: PageHeaderProps) {
    return (
        <div className="mb-8">
            <h1 className="text-4xl md:text-5xl font-bold text-foreground mb-2">{title}</h1>
            {subtitle && <p className="text-lg text-muted-foreground">{subtitle}</p>}
        </div>
    );
}
