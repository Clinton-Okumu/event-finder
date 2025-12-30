import { Button } from "@/components/ui/button";
import logo from "@/public/event.png";
import Image from "next/image";
import Link from "next/link";

const links = [
  { id: 1, title: "Home", href: "/" },
  { id: 2, title: "About us", href: "/about" },
  { id: 3, title: "Events", href: "/events" },
  { id: 4, title: "Testimonials", href: "/testimonials" },
  { id: 5, title: "Contactus", href: "/contact" },
];

export default function Navbar() {
  return (
    <header className="shadow-md sticky top-0 z-50 w-full border-b border-border bg-background">
      <nav className="mx-auto flex items-center px-4 py-4">
        <div className="flex items-center gap-3">
          <Image
            src={logo}
            alt="logo"
            width={40}
            height={40}
            className="rounded-md"
            priority
          />
          <span className="font-semibold text-lg tracking-tight text-foreground">
            Event-<span className="text-primary">Finder</span>
          </span>
        </div>

        <ul className="hidden md:flex items-center justify-center gap-6 flex-1">
          {links.map((l) => (
            <li key={l.id}>
              <Link
                href={l.href}
                className="font-medium transition-colors text-foreground/80 hover:text-primary"
              >
                {l.title}
              </Link>
            </li>
          ))}
        </ul>

        <div className="hidden md:flex items-center gap-3">
          <Button className="shadow-md">Web app</Button>
          <Button className="shadow-lg">Download app</Button>
        </div>
      </nav>
    </header>
  );
}
