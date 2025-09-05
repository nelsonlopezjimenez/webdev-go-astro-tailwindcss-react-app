---
import Layout from '../../layouts/Layout.astro';

export const prerender = false;

export async function getStaticPaths() {
  return [
    {params: {section: "section1-html-css"}},
    {params: {section: "section2-javascript"}},
    {params: {section: "section3-backend"}},
    {params: {section: "section4-react"}},
  ];
}

const { section } = Astro.params;

// Fetch section data
const API_BASE = 'http://localhost:8080/api';
let sectionData = null;
let error = null;

try {
  const response = await fetch(`${API_BASE}/sections/${section}`);
  if (response.ok) {
    sectionData = await response.json();
  } else if (response.status === 404) {
    error = 'Section not found';
  } else {
    error = 'Failed to load section';
  }
} catch (err) {
  console.error('Failed to fetch section:', err);
  error = 'Failed to connect to server';
}

// Section metadata
const sectionMeta = {
  "section1-html-css": {
    color: "blue",
    icon: "🌐",
    description: "Learn the foundation of web development with HTML structure and CSS styling",
    skills: ["HTML5 Elements", "CSS Flexbox & Grid", "Responsive Design", "Web Accessibility"]
  },
  "section2-javascript": {
    color: "green", 
    icon: "⚡",
    description: "Master JavaScript programming and DOM manipulation for interactive websites",
    skills: ["JavaScript ES6+", "DOM Manipulation", "Event Handling", "Async Programming"]
  },
  "section3-backend": {
    color: "purple",
    icon: "🚀",
    description: "Build server-side applications with Node.js and database integration",
    skills: ["Node.js & Express", "Database Design", "REST APIs", "Authentication"]
  },
  "section4-react": {
    color: "orange",
    icon: "⚛️",
    description: "Create modern web applications with React framework and advanced concepts",
    skills: ["React Components", "State Management", "Routing", "Deployment"]
  }
};

const meta = sectionMeta[section] || sectionMeta["section1-html-css"];
---

<Layout title={sectionData ? `${sectionData.name} | Edmonds College` : 'Section Not Found'}>
  <div class="min-h-screen bg-gray-50">
    
    {error ? (
      <!-- Error State -->
      <div class="max-w-4xl mx-auto px-6 py-16">
        <div class="bg-red-50 border border-red-200 rounded-xl p-12 text-center">
          <div class="text-6xl mb-6">❌</div>
          <h1 class="text-3xl font-bold text-red-800 mb-4">Section Not Available</h1>
          <p class="text-red-700 mb-8">{error}</p>
          <a 
            href="/" 
            class="inline-block bg-blue-600 text-white px-6 py-3 rounded-lg hover:bg-blue-700 transition-colors"
          >
            Back to Home
          </a>
        </div>
      </div>
    ) : sectionData ? (
      <>
        <!-- Section Header -->
        <div class={`bg-gradient-to-r from-${meta.color}-600 to-${meta.color}-700 text-white`}>
          <div class="max-w-7xl mx-auto px-6 py-16">
            
            <!-- Breadcrumb -->
            <nav class="mb-6">
              <ol class="flex items-center space-x-2 text-sm">
                <li><a href="/" class="text-white/80 hover:text-white">Home</a></li>
                <li class="text-white/60">/</li>
                <li class="text-white font-medium">{sectionData.name}</li>
              </ol>
            </nav>
            
            <div class="grid md:grid-cols-3 gap-8 items-center">
              <div class="md:col-span-2">
                <div class="flex items-center mb-4">
                  <span class="text-5xl mr-4">{meta.icon}</span>
                  <div>
                    <h1 class="text-4xl md:text-5xl font-bold mb-2">{sectionData.name}</h1>
                    <p class="text-xl text-white/90">{sectionData.description}</p>
                  </div>
                </div>
                
                <p class="text-lg text-white/80 mb-6 max-w-2xl">
                  {meta.description}
                </p>
                
                <div class="flex flex-wrap gap-2">
                  {meta.skills.map(skill => (
                    <span class="bg-white/20 backdrop-blur-sm text-white px-3 py-1 rounded-full text-sm font-medium">
                      {skill}
                    </span>
                  ))}
                </div>
              </div>
              
              <!-- Progress Card -->
              <div class="bg-white/10 backdrop-blur-sm rounded-xl p-6">
                <h3 class="text-lg font-bold mb-4">Section Progress</h3>
                <div class="space-y-4">
                  <div class="flex justify-between text-sm">
                    <span>Lessons Available</span>
                    <span>{sectionData.lessons.length} of 12</span>
                  </div>
                  <div class="w-full bg-white/20 rounded-full h-3">
                    <div 
                      class="bg-white h-3 rounded-full transition-all duration-500"
                      style={`width: ${(sectionData.lessons.length / 12) * 100}%`}
                    ></div>
                  </div>
                  
                  {sectionData.lessons.length > 0 && (
                    <div class="text-sm">
                      <p>Latest: Week {Math.max(...sectionData.lessons.map(l => l.week - sectionData.week_start + 1))}</p>
                      <p class="text-white/80 truncate">{sectionData.lessons[sectionData.lessons.length - 1]?.title}</p>
                    </div>
                  )}
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Lessons Grid -->
        <div class="max-w-7xl mx-auto px-6 py-16">
          {sectionData.lessons.length > 0 ? (
            <>
              <div class="mb-12">
                <h2 class="text-3xl font-bold text-gray-800 mb-4">Available Lessons</h2>
                <p class="text-lg text-gray-600">
                  Click on any lesson to start learning. Complete lessons in order for the best experience.
                </p>
              </div>
              
              <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
                {sectionData.lessons.map((lesson) => {
                  const weekInSection = lesson.week - sectionData.week_start + 1;
                  return (
                    <div class="bg-white rounded-xl shadow-lg hover:shadow-xl transition-all duration-300 border border-gray-200 overflow-hidden">
                      
                      <!-- Lesson Header -->
                      <div class={`bg-gradient-to-r from-${meta.color}-500 to-${meta.color}-600 text-white p-4`}>
                        <div class="flex items-center justify-between">
                          <span class="bg-white/20 backdrop-blur-sm px-3 py-1 rounded-full text-sm font-medium">
                            Week {weekInSection}
                          </span>
                          <span class="text-2xl">{meta.icon}</span>
                        </div>
                      </div>
                      
                      <!-- Lesson Content -->
                      <div class="p-6">
                        <h3 class="text-xl font-bold text-gray-800 mb-3 line-clamp-2">
                          {lesson.title}
                        </h3>
                        
                        {lesson.description && (
                          <p class="text-gray-600 mb-4 line-clamp-3">
                            {lesson.description}
                          </p>
                        )}
                        
                        <div class="flex items-center justify-between text-sm text-gray-500 mb-4">
                          <span>Global Week {lesson.week}</span>
                          <span>{new Date(lesson.created_at).toLocaleDateString()}</span>
                        </div>
                        
                        <a 
                          href={`/lessons/${lesson.week}`}
                          class={`block w-full text-center bg-${meta.color}-600 text-white px-4 py-3 rounded-lg hover:bg-${meta.color}-700 transform hover:scale-105 transition-all duration-200 font-medium`}
                        >
                          Start Lesson
                        </a>
                      </div>
                    </div>
                  );
                })}
              </div>
            </>
          ) : (
            <!-- No Lessons Available -->
            <div class="text-center py-16">
              <div class="text-8xl mb-6">📚</div>
              <h2 class="text-3xl font-bold text-gray-800 mb-4">Lessons Coming Soon</h2>
              <p class="text-xl text-gray-600 mb-8 max-w-2xl mx-auto">
                This section is being prepared with amazing content. Check back soon for new lessons!
              </p>
              
              <div class="bg-yellow-50 border border-yellow-200 rounded-xl p-6 max-w-md mx-auto">
                <h3 class="font-bold text-yellow-800 mb-3">Expected Content:</h3>
                <ul class="text-sm text-yellow-700 space-y-1">
                  {meta.skills.map(skill => (
                    <li class="flex items-center">
                      <svg class="w-4 h-4 mr-2 text-yellow-600" fill="currentColor" viewBox="0 0 20 20">
                        <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd"></path>
                      </svg>
                      {skill}
                    </li>
                  ))}
                </ul>
              </div>
            </div>
          )}
        </div>

        <!-- Section Navigation -->
        <div class="bg-white border-t border-gray-200 py-12">
          <div class="max-w-7xl mx-auto px-6">
            <h3 class="text-2xl font-bold text-gray-800 mb-6 text-center">Explore Other Sections</h3>
            
            <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
              {Object.entries(sectionMeta).map(([sectionId, meta]) => {
                const isCurrentSection = sectionId === section;
                return (
                  <a 
                    href={isCurrentSection ? '#' : `/sections/${sectionId}`}
                    class={`p-4 rounded-lg text-center transition-all duration-200 ${
                      isCurrentSection 
                        ? 'bg-gray-800 text-white cursor-default' 
                        : `bg-${meta.color}-50 hover:bg-${meta.color}-100 border border-${meta.color}-200`
                    }`}
                  >
                    <div class="text-2xl mb-2">{meta.icon}</div>
                    <div class={`text-sm font-medium ${isCurrentSection ? 'text-white' : `text-${meta.color}-800`}`}>
                      Section {sectionId.slice(-1)}
                    </div>
                    {isCurrentSection && (
                      <div class="text-xs text-gray-300 mt-1">Current</div>
                    )}
                  </a>
                );
              })}
            </div>
          </div>
        </div>
      </>
    ) : (
      <!-- Loading State -->
      <div class="max-w-4xl mx-auto px-6 py-16">
        <div class="bg-white rounded-xl shadow-lg p-12 text-center">
          <div class="animate-spin w-16 h-16 border-4 border-blue-500 border-t-transparent rounded-full mx-auto mb-6"></div>
          <p class="text-xl text-gray-600">Loading section...</p>
        </div>
      </div>
    )}
  </div>
</Layout>

<style>
  /* Prevent Tailwind purging dynamic classes */
  .text-blue-800, .bg-blue-50, .bg-blue-100, .bg-blue-600, .bg-blue-700, .border-blue-200,
  .text-green-800, .bg-green-50, .bg-green-100, .bg-green-600, .bg-green-700, .border-green-200,
  .text-purple-800, .bg-purple-50, .bg-purple-100, .bg-purple-600, .bg-purple-700, .border-purple-200,
  .text-orange-800, .bg-orange-50, .bg-orange-100, .bg-orange-600, .bg-orange-700, .border-orange-200 {
    /* Ensure these classes are available */
  }
  
  .line-clamp-2 {
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
    overflow: hidden;
  }
  
  .line-clamp-3 {
    display: -webkit-box;
    -webkit-line-clamp: 3;
    -webkit-box-orient: vertical;
    overflow: hidden;
  }
</style>