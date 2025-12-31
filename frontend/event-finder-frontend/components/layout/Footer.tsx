"use client";
import { Button } from "@/components/ui/button";
import logo from "@/public/event.png";
import { Facebook, Instagram, Mail, Twitter } from "lucide-react";
import Image from "next/image";

const Footer = () => {
  return (
    <footer className="bg-muted text-foreground pt-12 pb-6 shadow-inner">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        {/* Top */}
        <div className="grid grid-cols-1 md:grid-cols-4 gap-8 border-b border-border pb-10">
          {/* Logo */}
          <div>
            <div className="flex items-center gap-3">
              <Image src={logo} alt="TourVista logo" width={40} height={40} />
              <span className="text-xl font-bold">
                Event-<span className="text-primary">Finder</span>
              </span>
            </div>

            <p className="mt-4 text-sm text-muted-foreground">
              Explore the world with confidence. Book your next adventure with
              Event-Finder.
            </p>
          </div>

          {/* Links */}
          <div>
            <h4 className="font-semibold mb-4">Company</h4>
            <ul className="space-y-2 text-sm">
              <li>
                <a className="hover:opacity-60" href="#">
                  About Us
                </a>
              </li>
              <li>
                <a className="hover:opacity-60" href="#">
                  Careers
                </a>
              </li>
              <li>
                <a className="hover:opacity-60" href="#">
                  Blog
                </a>
              </li>
            </ul>
          </div>

          {/* Support */}
          <div>
            <h4 className="font-semibold mb-4">Support</h4>
            <ul className="space-y-2 text-sm">
              <li>
                <a className="hover:opacity-60" href="#">
                  Help Center
                </a>
              </li>
              <li>
                <a className="hover:opacity-60" href="#">
                  Contact
                </a>
              </li>
              <li>
                <a className="hover:opacity-60" href="#">
                  Terms & Privacy
                </a>
              </li>
            </ul>
          </div>

          {/* Newsletter */}
          <div>
            <h4 className="font-semibold mb-4">Subscribe</h4>
            <p className="text-sm text-muted-foreground mb-4">
              Get updates on new tours and offers.
            </p>
            <form className="flex flex-col sm:flex-row gap-2">
              <input
                type="email"
                placeholder="Enter your email"
                className="px-4 py-2 rounded-md border border-border focus:outline-none focus:ring-1 focus:ring-primary"
              />
              <Button
                type="submit"
                className="px-4 py-2 rounded-md border border-border hover:bg-accent transition"
              >
                Subscribe
              </Button>
            </form>
          </div>
        </div>

        {/* Bottom */}
        <div className="flex flex-col sm:flex-row justify-between items-center mt-8 text-sm text-muted-foreground">
          <p>
            &copy; {new Date().getFullYear()} Event-Finder. All rights reserved.
          </p>

          <div className="flex gap-4 mt-4 sm:mt-0">
            <a className="hover:opacity-60" href="#">
              <Facebook size={20} />
            </a>
            <a className="hover:opacity-60" href="#">
              <Twitter size={20} />
            </a>
            <a className="hover:opacity-60" href="#">
              <Instagram size={20} />
            </a>
            <a className="hover:opacity-60" href="mailto:hello@tourvista.com">
              <Mail size={20} />
            </a>
          </div>
        </div>
      </div>
    </footer>
  );
};

export default Footer;
