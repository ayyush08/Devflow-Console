import { Hero } from "@/components/Hero";
import Navbar from "@/components/Navbar";


export default function Home() {
  return (
        <div className="relative flex flex-col items-center justify-center">
          <Navbar/>
          <Hero/>
        </div>
  );
}
