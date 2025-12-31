import ChatInterface from "@/components/support/ChatInterface";
import ContactOptions from "@/components/support/ContactOptions";
import FAQSection from "@/components/support/FAQSection";
import PageHeader from "@/components/layout/PageHeader";

export default function SupportPage() {
  return (
    <div className="container mx-auto px-4 py-8">
      <PageHeader
        title="Customer Support"
        subtitle="We're here to help you"
      />

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-8 mt-8">
        <div className="lg:col-span-2 space-y-8">
          <ChatInterface />
        </div>

        <div className="space-y-8">
          <ContactOptions />
          <FAQSection />
        </div>
      </div>
    </div>
  );
}
