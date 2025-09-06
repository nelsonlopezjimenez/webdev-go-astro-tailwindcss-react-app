---
import Layout from '../layouts/Layout.astro';

// Fetch websites data
const API_BASE = 'http://localhost:8080/api';
let websites = [];
let error = null;

try {
  const response = await fetch(`${API_BASE}/websites`);
  if (response.ok) {
    websites = await response.json();
  } else {
    error = 'Failed to load websites';
  }
} catch (err) {
  console.error('Failed to fetch websites:', err);
  error = 'Failed to connect to server';
}

// Helper function to format file size
function formatFileSize(bytes) {
  if (bytes === 0) return '0 Bytes';
  const k = 1024;
  const sizes = ['Bytes', 'KB', 'MB', 'GB'];
  const i = Math.floor(Math.log(bytes) / Math.log(k));
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
}

// Helper function to get domain from URL
function getDomain(url) {
  if (!url) return 'Unknown';
  try {
    return new URL(url).hostname;
  } catch {
    return 'Unknown';
  }
}
---

<Layout title="Offline Websites | Edmonds College">
  <div class="min-h-screen bg-gray-50">
    
    <!-- Header -->
    <div class="bg-gradient-to-r from-purple-600 to-blue-600 text-white">
      <div class="max-w-7xl mx-auto px-6 py-16">
        <nav class="mb-6">
          <ol class="flex items-center space-x-2 text-sm">
            <li><a href="/" class="text-white opacity-80 hover:opacity-100">Home</a></li>
            <li class="text-white opacity-60">/</li>
            <li class="text-white font-medium">Offline Websites</li>
          </ol>
        </nav>
        
        <div class="flex items-center mb-6">
          <span class="text-5xl mr-4">🌐</span>
          <div>
            <h1 class="text-4xl md:text-5xl font-bold mb-2">Offline Websites</h1>
            <p class="text-xl text-white opacity-90">HTTrack Downloaded Sites</p>
          </div>
        </div>
        
        <p class="text-lg text-white opacity-80 max-w-3xl">
          Access downloaded websites for offline study and reference. These sites have been archived using HTTrack and are available even without internet connection.
        </p>
        
        <div class="mt-6 bg-white bg-opacity-20 backdrop-blur-sm rounded-lg px-4 py-3 inline-block">
          <div class="flex items-center space-x-4 text-sm">
            <span>Total Sites: <strong>{websites.length}</strong></span>
            <span>•</span>
            <span>Total Size: <strong>{formatFileSize(websites.reduce((sum, site) => sum + (site.size || 0), 0))}</strong></span>
          </div>
        </div>
      </div>
    </div>

    <!-- Main Content -->
    <div class="max-w-7xl mx-auto px-6 py-12">
      
      {error ? (
        /* Error State */
        <div class="bg-red-50 border border-red-200 rounded-xl p-12 text-center">
          <div class="text-6xl mb-6">❌</div>
          <h2 class="text-3xl font-bold text-red-800 mb-4">Unable to Load Websites</h2>
          <p class="text-red-700 mb-8">{error}</p>
          <div class="bg-red-100 border border-red-300 rounded-lg p-4 max-w-md mx-auto">
            <h3 class="font-bold text-red-800 mb-2">Possible Issues:</h3>
            <ul class="text-sm text-red-700 text-left space-y-1">
              <li>• Websites directory not found at C:\websites</li>
              <li>• No HTTrack websites in the directory</li>
              <li>• Server permission issues</li>
              <li>• Server not properly configured</li>
            </ul>
          </div>
        </div>
      ) : websites.length === 0 ? (
        /* No Websites State */
        <div class="bg-yellow-50 border border-yellow-200 rounded-xl p-12 text-center">
          <div class="text-8xl mb-6">📂</div>
          <h2 class="text-3xl font-bold text-yellow-800 mb-4">No Websites Found</h2>
          <p class="text-xl text-yellow-700 mb-8">
            No HTTrack websites were found in the configured directory.
          </p>
          <div class="bg-yellow-100 border border-yellow-300 rounded-lg p-6 max-w-2xl mx-auto">
            <h3 class="font-bold text-yellow-800 mb-3">To add websites:</h3>
            <ol class="text-sm text-yellow-700 space-y-2 text-left">
              <li class="flex items-start">
                <span class="w-6 h-6 bg-yellow-500 text-white rounded-full flex items-center justify-center text-xs font-bold mr-3 mt-0.5">1</span>
                Use HTTrack to download websites to C:\websites
              </li>
              <li class="flex items-start">
                <span class="w-6 h-6 bg-yellow-500 text-white rounded-full flex items-center justify-center text-xs font-bold mr-3 mt-0.5">2</span>
                Ensure each website has an index.html file
              </li>
              <li class="flex items-start">
                <span class="w-6 h-6 bg-yellow-500 text-white rounded-full flex items-center justify-center text-xs font-bold mr-3 mt-0.5">3</span>
                Restart the server to scan for new websites
              </li>
            </ol>
          </div>
        </div>
      ) : (
        /* Websites Grid */
        <div>
          <div class="flex items-center justify-between mb-8">
            <h2 class="text-2xl font-bold text-gray-800">Available Websites</h2>
            <div class="text-sm text-gray-500">
              Click any website to open in a new tab
            </div>
          </div>
          
          <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
            {websites.map((website) => {
              const domain = getDomain(website.url);
              const lastScanned = new Date(website.last_scanned).toLocaleDateString();
              
              return (
                <div class="bg-white rounded-xl shadow-lg hover:shadow-xl transition-all duration-300 border border-gray-200 overflow-hidden group">
                  
                  {/* Website Header */}
                  <div class="bg-gradient-to-r from-purple-500 to-blue-500 text-white p-4">
                    <div class="flex items-center justify-between">
                      <div class="flex items-center space-x-2">
                        <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
                          <path fill-rule="evenodd" d="M4.083 9h1.946c.089-1.546.383-2.97.837-4.118C6.004 2.42 4.674 1 3 1a1 1 0 000 2c.83 0 1.542.631 1.83 1.495A11.86 11.86 0 004.083 9zM10 2a1 1 0 01-1 1c-.83 0-1.542.631-1.83 1.495A11.86 11.86 0 006.417 9H7.5a9 9 0 011.5-5.5c1.5 0 2.5 2 2.5 5.5h1.083c-.089-1.546-.383-2.97-.837-4.118C13.496 2.42 14.826 1 16.5 1a1 1 0 000 2c-.83 0-1.542.631-1.83 1.495A11.86 11.86 0 0015.917 9z" clip-rule="evenodd"></path>
                        </svg>
                        <span class="font-medium text-sm">Offline Site</span>
                      </div>
                      <div class="bg-white bg-opacity-20 backdrop-blur-sm px-2 py-1 rounded text-xs">
                        {formatFileSize(website.size)}
                      </div>
                    </div>
                  </div>
                  
                  {/* Website Content */}
                  <div class="p-6">
                    <h3 class="text-xl font-bold text-gray-800 mb-3 group-hover:text-purple-600 transition-colors">
                      {website.name}
                    </h3>
                    
                    <div class="space-y-2 mb-4">
                      {website.url && (
                        <div class="flex items-center text-sm text-gray-600">
                          <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.102m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1"></path>
                          </svg>
                          <span class="truncate">{domain}</span>
                        </div>
                      )}
                      
                      <div class="flex items-center text-sm text-gray-600">
                        <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"></path>
                        </svg>
                        <span>Scanned: {lastScanned}</span>
                      </div>
                      
                      <div class="flex items-center text-sm text-gray-600">
                        <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                        </svg>
                        <span>Index: {website.index_file}</span>
                      </div>
                    </div>
                    
                    {/* Action Buttons */}
                    <div class="space-y-3">
                      <a 
                        href={`/websites/${website.name}`}
                        target="_blank"
                        rel="noopener noreferrer"
                        class="block w-full text-center bg-purple-600 text-white px-4 py-3 rounded-lg hover:bg-purple-700 transform hover:scale-105 transition-all duration-200 font-medium"
                      >
                        Open Website
                      </a>
                      
                      <div class="grid grid-cols-2 gap-2">
                        <a 
                          href={`/websites/${website.name}/${website.index_file}`}
                          target="_blank"
                          rel="noopener noreferrer"
                          class="text-center bg-gray-100 text-gray-700 px-3 py-2 rounded hover:bg-gray-200 transition-colors text-sm"
                        >
                          Direct Link
                        </a>
                        {website.url && (
                          <a 
                            href={website.url}
                            target="_blank"
                            rel="noopener noreferrer"
                            class="text-center bg-blue-100 text-blue-700 px-3 py-2 rounded hover:bg-blue-200 transition-colors text-sm"
                          >
                            Original URL
                          </a>
                        )}
                      </div>
                    </div>
                  </div>
                </div>
              );
            })}
          </div>
          
          {/* Usage Instructions */}
          <div class="mt-12 bg-blue-50 border border-blue-200 rounded-xl p-8">
            <h3 class="text-xl font-bold text-blue-800 mb-4">Using Offline Websites</h3>
            <div class="grid md:grid-cols-2 gap-6">
              <div>
                <h4 class="font-semibold text-blue-700 mb-2">Features</h4>
                <ul class="text-blue-700 text-sm space-y-1">
                  <li>• Browse websites without internet connection</li>
                  <li>• All linked pages and resources included</li>
                  <li>• Original formatting and functionality preserved</li>
                  <li>• Fast local loading times</li>
                </ul>
              </div>
              <div>
                <h4 class="font-semibold text-blue-700 mb-2">Tips</h4>
                <ul class="text-blue-700 text-sm space-y-1">
                  <li>• Use "Open Website" for normal browsing</li>
                  <li>• "Direct Link" opens the index page directly</li>
                  <li>• "Original URL" shows the live website (requires internet)</li>
                  <li>• All internal links within the site work offline</li>
                </ul>
              </div>
            </div>
          </div>
        </div>
      )}
    </div>
  </div>
</Layout>