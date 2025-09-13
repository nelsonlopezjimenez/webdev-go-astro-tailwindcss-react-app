// src/components/DynamicTableOfContents.jsx
import { useState, useEffect } from 'react';

export default function DynamicTableOfContents({ sectionName, weekNumber }) {
  const [tocItems, setTocItems] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchTOC();
  }, [sectionName, weekNumber]);

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

  return (
    <nav className="bg-white rounded-lg shadow-sm border border-gray-200 p-4">
      <h3 className="font-semibold mb-4">Week {weekNumber} Contents</h3>
      <ul className="space-y-2">
        {tocItems.map((item, index) => (
          <li key={index}>
            <a href={`#${item.id}`} className="text-sm text-gray-700 hover:text-indigo-600">
              {item.title}
            </a>
          </li>
        ))}
      </ul>
    </nav>
  );
}