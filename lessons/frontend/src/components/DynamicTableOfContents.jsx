import { useState, useEffect } from 'react';

export default function DynamicTableOfContents({ sectionName, weekNumber }) {
  const [tocItems, setTocItems] = useState([]);
  const [activeId, setActiveId] = useState('');
  const [loading, setLoading] = useState(true);

  // Your existing fetchTOC code...

  useEffect(() => {
    const handleScroll = () => {
      const headings = document.querySelectorAll('h1[id], h2[id], h3[id], h4[id], h5[id], h6[id]');
      let currentActiveId = '';
      
      for (const heading of headings) {
        const rect = heading.getBoundingClientRect();
        if (rect.top <= 150) {
          currentActiveId = heading.id;
        }
      }
      
      setActiveId(currentActiveId);
    };

    window.addEventListener('scroll', handleScroll, { passive: true });
    handleScroll(); // Set initial active state
    
    return () => window.removeEventListener('scroll', handleScroll);
  }, []);

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
                  ? 'text-indigo-600 font-semibold border-l-2 border-indigo-600 pl-2' 
                  : 'text-gray-700 hover:text-indigo-600 pl-2'
              }`}
              onClick={(e) => {
                e.preventDefault();
                const element = document.getElementById(item.id);
                if (element) {
                  element.scrollIntoView({ behavior: 'smooth', block: 'start' });
                  setActiveId(item.id);
                }
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