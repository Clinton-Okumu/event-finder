import AboutHero from "@/components/about/AboutHero";
import AboutMission from "@/components/about/AboutMission";
import AboutStats from "@/components/about/AboutStats";
import AboutStory from "@/components/about/AboutStory";
import AboutTeam from "@/components/about/AboutTeam";
import PageHeader from "@/components/layout/PageHeader";

export default function AboutPage() {
  return (
    <div className="container mx-auto px-4 py-8">
      <PageHeader title="About Us" subtitle="Learn more about Event Finder" />
      <AboutHero />
      <AboutMission />
      <AboutStory />
      <AboutStats />
      <AboutTeam />
    </div>
  );
}
