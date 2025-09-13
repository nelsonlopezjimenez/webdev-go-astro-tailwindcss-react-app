// src/components/DynamicTableOfContents.jsx
import { useState, useEffect } from 'react';

export default function DynamicTableOfContents({ sectionName, weekNumber }) {
  const [tocItems, setTocItems] = useState([]);
  const [loading, setLoading] = useState(true);
   const [activeId, setActiveId] = useState('');

  useEffect(() => {
    fetchTOC();
  }, [sectionName, weekNumber]);


   // NEW: Add this useEffect for Intersection Observer
  useEffect(() => {
    const observer = new IntersectionObserver(
      (entries) => {
        entries.forEach(entry => {
          const id = entry.target.id;
          
          if (entry.isIntersecting) {
            setActiveId(id);
          }
        });
      },
      { 
        rootMargin: '-20% 0px -80% 0px',
        threshold: 0
      }
    );

    // Wait a bit for content to load, then observe headings
    const timer = setTimeout(() => {
      document.querySelectorAll('h1[id], h2[id], h3[id], h4[id], h5[id], h6[id]').forEach(
        heading => observer.observe(heading)
      );
    }, 1000);

    return () => {
      clearTimeout(timer);
      observer.disconnect();
    };
  }, [tocItems]); // Re-run when tocItems change


  const fetchTOC = async () => {
    setLoading(true);
    try {
      const response = await fetch(`/api/sections/${sectionName}/week/${weekNumber}/toc`);
      
      if (response.ok) {
        const data = await response.json();
        setTocItems(data.tocItems || []);
      } else {
        setTocItems(getDefaultTOCItems());
      }
    } catch (error) {
      console.error('Failed to fetch TOC:', error);
      setTocItems(getDefaultTOCItems());
    } finally {
      setLoading(false);
    }
  };

  const getDefaultTOCItems = () => [
    { id: "learning-objectives", title: "Learning Objectives", level: 2 },
    { id: "introduction", title: "Introduction", level: 2 }
  ];

  if (loading) {
    return <div className="p-4">Loading TOC...</div>;
  }

 // Update your render to use activeId state
  return (
    <nav className="bg-white rounded-lg shadow-sm border border-gray-200 p-4 sticky top-4">
      <h3 className="font-semibold mb-4">Week {weekNumber} Contents</h3>
      <ul className="space-y-2">
        {tocItems.map((item, index) => (
          <li key={index}>
            <a 
              href={`#${item.id}`} 
              className={`block text-sm py-1 transition-colors ${
                activeId === item.id 
                  ? 'text-indigo-600 font-semibold' 
                  : 'text-gray-700 hover:text-indigo-600'
              }`}
              onClick={(e) => {
                e.preventDefault();
                document.getElementById(item.id)?.scrollIntoView({ 
                  behavior: 'smooth',
                  block: 'start'
                });
              }}
            >
              {item.title}
            </a>
          </li>
        ))}
      </ul>
    </nav>
  );
}
