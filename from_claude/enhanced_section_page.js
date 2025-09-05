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

// Fetch section data and syllabus
const API_BASE = 'http://localhost:8080/api';
let sectionData = null;
let syllabusData = null;
let error = null;

try {
  const [sectionRes, syllabusRes] = await Promise.all([
    fetch(`${API_BASE}/sections/${section}`),
    fetch(`${API_BASE}/sections/${section}/syllabus`)
  ]);
  
  if (sectionRes.ok) {
    sectionData = await sectionRes.json();
  }
  
  if (syllabusRes.ok) {
    syllabusData = await syllabusRes.json();
  } else {
    // Fallback to basic section data
    syllabusData = sectionData;
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
    quarter: "Fall Quarter"
  },
  "section2-javascript": {
    color: "green", 
    icon: "⚡",
    description: "Master JavaScript programming and DOM manipulation for interactive websites",
    quarter: "Winter Quarter"
  },
  "section3-backend": {
    color: "purple",
    icon: "🚀",
    description: "Build server-side applications with Node.js and database integration",
    quarter: "Spring Quarter"
  },
  "section4-react": {
    color: "orange",
    icon: "⚛️",
    description: "Create modern web applications with React framework and advanced concepts",
    quarter: "Summer Quarter"
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
                    <div class="flex items-center gap-3 mb-2">
                      <h1 class="text-4xl md:text-5xl font-bold">{sectionData.name}</h1>
                      {syllabusData?.syllabus_info?.course_code && (
                        <span class="bg-white/20 backdrop-blur-sm px-3 py-1 rounded-full text-sm font-medium">
                          {syllabusData.syllabus_info.course_code}
                        </span>
                      )}
                    </div>
                    <p class="text-xl text-white/90">{meta.quarter}</p>
                  </div>
                </div>
                
                <p class="text-lg text-white/80 mb-6 max-w-2xl">
                  {meta.description}
                </p>
                
                {syllabusData?.syllabus_info?.credits && (
                  <div class="flex items-center gap-4 text-white/90">
                    <span class="flex items-center">
                      <svg class="w-4 h-4 mr-2" fill="currentColor" viewBox="0 0 20 20">
                        <path d="M10.394 2.08a1 1 0 00-.788 0l-7 3a1 1 0 000 1.84L5.25 8.051a.999.999 0 01.356-.257l4-1.714a1 1 0 11.788 1.838L7.667 9.088l1.94.831a1 1 0 00.787 0l7-3a1 1 0 000-1.838l-7-3z"></path>
                      </svg>
                      {syllabusData.syllabus_info.credits}
                    </span>
                    {syllabusData.syllabus_info.prerequisites && (
                      <span class="text-sm">
                        Prerequisite: {syllabusData.syllabus_info.prerequisites}
                      </span>
                    )}
                  </div>
                )}
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

        <!-- Tab Navigation -->
        <div class="bg-white shadow-sm sticky top-0 z-10">
          <div class="max-w-7xl mx-auto px-6">
            <nav class="flex space-x-8" id="tab-navigation">
              <button 
                class="tab-button py-4 px-2 border-b-2 font-medium text-sm focus:outline-none transition-colors active"
                data-tab="lessons"
              >
                Lessons ({sectionData.lessons.length})
              </button>
              <button 
                class="tab-button py-4 px-2 border-b-2 font-medium text-sm focus:outline-none transition-colors"
                data-tab="syllabus"
              >
                Syllabus
              </button>
              <button 
                class="tab-button py-4 px-2 border-b-2 font-medium text-sm focus:outline-none transition-colors"
                data-tab="schedule"
              >
                Weekly Schedule
              </button>
            </nav>
          </div>
        </div>

        <!-- Tab Content -->
        <div class="max-w-7xl mx-auto px-6 py-8">
          
          <!-- Lessons Tab -->
          <div id="lessons-tab" class="tab-content">
            {sectionData.lessons.length > 0 ? (
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
            ) : (
              <div class="text-center py-16">
                <div class="text-8xl mb-6">📚</div>
                <h2 class="text-3xl font-bold text-gray-800 mb-4">Lessons Coming Soon</h2>
                <p class="text-xl text-gray-600 mb-8 max-w-2xl mx-auto">
                  This section is being prepared with amazing content. Check back soon for new lessons!
                </p>
              </div>
            )}
          </div>

          <!-- Syllabus Tab -->
          <div id="syllabus-tab" class="tab-content hidden">
            {syllabusData?.syllabus_info ? (
              <div class="grid lg:grid-cols-3 gap-8">
                <!-- Course Overview -->
                <div class="lg:col-span-2 space-y-8">
                  
                  <!-- Course Description -->
                  <div class="bg-white rounded-xl shadow-lg p-8">
                    <h2 class="text-2xl font-bold text-gray-800 mb-4">Course Description</h2>
                    <p class="text-gray-700 leading-relaxed">
                      {syllabusData.syllabus_info.description}
                    </p>
                  </div>

                  <!-- Learning Objectives -->
                  <div class="bg-white rounded-xl shadow-lg p-8">
                    <h2 class="text-2xl font-bold text-gray-800 mb-6">Learning Objectives</h2>
                    <p class="text-gray-600 mb-4">Upon successful completion of this course, students will be able to:</p>
                    <ul class="space-y-3">
                      {syllabusData.syllabus_info.objectives?.map((objective, index) => (
                        <li class="flex items-start">
                          <span class={`bg-${meta.color}-500 text-white rounded-full w-6 h-6 flex items-center justify-center text-sm font-bold mr-3 mt-0.5 flex-shrink-0`}>
                            {index + 1}
                          </span>
                          <span class="text-gray-700">{objective}</span>
                        </li>
                      ))}
                    </ul>
                  </div>

                  <!-- Course Topics -->
                  <div class="bg-white rounded-xl shadow-lg p-8">
                    <h2 class="text-2xl font-bold text-gray-800 mb-6">Course Topics</h2>
                    <div class="grid md:grid-cols-2 gap-4">
                      {syllabusData.syllabus_info.topics?.map((topic, index) => (
                        <div class="flex items-center p-3 bg-gray-50 rounded-lg">
                          <div class={`w-2 h-2 bg-${meta.color}-500 rounded-full mr-3 flex-shrink-0`}></div>
                          <span class="text-gray-700 text-sm">{topic}</span>
                        </div>
                      ))}
                    </div>
                  </div>
                </div>

                <!-- Sidebar Information -->
                <div class="space-y-6">
                  
                  <!-- Course Details -->
                  <div class="bg-white rounded-xl shadow-lg p-6">
                    <h3 class="text-lg font-bold text-gray-800 mb-4">Course Details</h3>
                    <div class="space-y-3">
                      <div class="flex justify-between">
                        <span class="text-gray-600">Course Code</span>
                        <span class="font-semibold">{syllabusData.syllabus_info.course_code}</span>
                      </div>
                      <div class="flex justify-between">
                        <span class="text-gray-600">Credits</span>
                        <span class="font-semibold">{syllabusData.syllabus_info.credits}</span>
                      </div>
                      <div class="flex justify-between">
                        <span class="text-gray-600">Quarter</span>
                        <span class="font-semibold">{meta.quarter}</span>
                      </div>
                      <div class="pt-3 border-t border-gray-200">
                        <span class="text-gray-600 text-sm">Prerequisites</span>
                        <p class="text-sm text-gray-800 mt-1">{syllabusData.syllabus_info.prerequisites}</p>
                      </div>
                    </div>
                  </div>

                  <!-- Assessment -->
                  <div class="bg-white rounded-xl shadow-lg p-6">
                    <h3 class="text-lg font-bold text-gray-800 mb-4">Assessment</h3>
                    <div class="space-y-3">
                      {syllabusData.syllabus_info.assessment?.map((item, index) => (
                        <div class="text-sm">
                          <div class="flex items-center justify-between">
                            <span class="text-gray-700">{item}</span>
                          </div>
                        </div>
                      ))}
                    </div>
                  </div>

                  <!-- Resources -->
                  <div class="bg-white rounded-xl shadow-lg p-6">
                    <h3 class="text-lg font-bold text-gray-800 mb-4">Resources</h3>
                    <div class="space-y-2">
                      {syllabusData.syllabus_info.resources?.map((resource, index) => (
                        <div class="flex items-start text-sm">
                          <div class={`w-2 h-2 bg-${meta.color}-500 rounded-full mr-2 mt-2 flex-shrink-0`}></div>
                          <span class="text-gray-700">{resource}</span>
                        </div>
                      ))}
                    </div>
                  </div>
                </div>
              </div>
            ) : (
              <div class="text-center py-16">
                <div class="text-6xl mb-4">📋</div>
                <h2 class="text-2xl font-bold text-gray-800 mb-4">Syllabus Information Coming Soon</h2>
                <p class="text-gray-600">Detailed syllabus will be available closer to the start date.</p>
              </div>
            )}
          </div>

          <!-- Schedule Tab -->
          <div id="schedule-tab" class="tab-content hidden">
            <div class="bg-white rounded-xl shadow-lg p-8">
              <h2 class="text-2xl font-bold text-gray-800 mb-6">Weekly Schedule</h2>
              
              <div class="space-y-4">
                {Array.from({length: 12}, (_, i) => {
                  const weekNumber = i + 1;
                  const globalWeek = sectionData.week_start + i;
                  const lesson = sectionData.lessons.find(l => l.week === globalWeek);
                  const isAvailable = !!lesson;
                  
                  return (
                    <div class={`border rounded-lg p-6 transition-all ${
                      isAvailable 
                        ? `border-${meta.color}-200 bg-${meta.color}-50` 
                        : 'border-gray-200 bg-gray-50'
                    }`}>
                      <div class="flex items-center justify-between">
                        <div class="flex-1">
                          <div class="flex items-center gap-4 mb-2">
                            <span class={`font-bold text-lg ${
                              isAvailable ? `text-${meta.color}-800` : 'text-gray-500'
                            }`}>
                              Week {weekNumber}
                            </span>
                            <span class="text-sm text-gray-500">
                              (Global Week {globalWeek})
                            </span>
                            {isAvailable && (
                              <span class={`bg-${meta.color}-500 text-white px-2 py-1 rounded-full text-xs font-medium`}>
                                Available
                              </span>
                            )}
                          </div>
                          
                          <h3 class={`text-lg font-semibold mb-2 ${
                            isAvailable ? 'text-gray-800' : 'text-gray-500'
                          }`}>
                            {lesson?.title || `Week ${weekNumber} - Coming Soon`}
                          </h3>
                          
                          {lesson?.description && (
                            <p class="text-gray-600 mb-3">{lesson.description}</p>
                          )}
                          
                          {isAvailable && lesson.created_at && (
                            <p class="text-sm text-gray-500">
                              Published: {new Date(lesson.created_at).toLocaleDateString()}
                            </p>
                          )}
                        </div>
                        
                        <div class="ml-6">
                          {isAvailable ? (
                            <a 
                              href={`/lessons/${globalWeek}`}
                              class={`bg-${meta.color}-600 text-white px-6 py-2 rounded-lg hover:bg-${meta.color}-700 transition-colors font-medium`}
                            >
                              View Lesson
                            </a>
                          ) : (
                            <div class="bg-gray-300 text-gray-500 px-6 py-2 rounded-lg cursor-not-allowed font-medium">
                              Not Available
                            </div>
                          )}
                        </div>
                      </div>
                    </div>
                  );
                })}
              </div>
            </div>
          </div>
        </div>
      </>
    )}
  </div>
</Layout>

<script>
  // Tab functionality
  document.addEventListener('DOMContentLoaded', function() {
    const tabButtons = document.querySelectorAll('.tab-button');
    const tabContents = document.querySelectorAll('.tab-content');
    
    function switchTab(targetTab) {
      // Remove active class from all buttons
      tabButtons.forEach(button => {
        button.classList.remove('active', 'border-blue-500', 'text-blue-600');
        button.classList.add('border-transparent', 'text-gray-500', 'hover:text-gray-700', 'hover:border-gray-300');
      });
      
      // Hide all tab contents
      tabContents.forEach(content => {
        content.classList.add('hidden');
      });
      
      // Show target tab content
      const targetContent = document.getElementById(targetTab + '-tab');
      if (targetContent) {
        targetContent.classList.remove('hidden');
      }
      
      // Activate clicked button
      const targetButton = document.querySelector(`[data-tab="${targetTab}"]`);
      if (targetButton) {
        targetButton.classList.add('active', 'border-blue-500', 'text-blue-600');
        targetButton.classList.remove('border-transparent', 'text-gray-500', 'hover:text-gray-700', 'hover:border-gray-300');
      }
    }
    
    // Add click listeners to tab buttons
    tabButtons.forEach(button => {
      button.addEventListener('click', function() {
        const tabName = this.getAttribute('data-tab');
        switchTab(tabName);
      });
    });
    
    // Initialize first tab as active
    switchTab('lessons');
  });
</script>

<style>
  /* Prevent Tailwind purging dynamic classes */
  .text-blue-800, .bg-blue-50, .bg-blue-100, .bg-blue-600, .bg-blue-700, .border-blue-200, .from-blue-500, .to-blue-600, .from-blue-600, .to-blue-700,
  .text-green-800, .bg-green-50, .bg-green-100, .bg-green-600, .bg-green-700, .border-green-200, .from-green-500, .to-green-600, .from-green-600, .to-green-700,
  .text-purple-800, .bg-purple-50, .bg-purple-100, .bg-purple-600, .bg-purple-700, .border-purple-200, .from-purple-500, .to-purple-600, .from-purple-600, .to-purple-700,
  .text-orange-800, .bg-orange-50, .bg-orange-100, .bg-orange-600, .bg-orange-700, .border-orange-200, .from-orange-500, .to-orange-600, .from-orange-600, .to-orange-700 {
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
  
  .tab-button.active {
    border-color: rgb(59 130 246);
    color: rgb(37 99 235);
  }
  
  .tab-button:not(.active) {
    border-color: transparent;
    color: rgb(107 114 128);
  }
  
  .tab-button:not(.active):hover {
    color: rgb(55 65 81);
    border-color: rgb(209 213 219);
  }
</style>