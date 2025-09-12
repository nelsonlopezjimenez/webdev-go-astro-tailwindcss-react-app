import React, { useState, useEffect } from 'react';
import { marked } from 'marked';

const MarkdownRenderer = ({ content }) => {
  const [htmlContent, setHtmlContent] = useState('');

  useEffect(() => {
    if (content) {
      console.log("line 9 ", typeof content)
      // Configure marked options
      marked.setOptions({
        breaks: true,
        gfm: true,
        highlight: function(code, lang) {
          // Basic syntax highlighting placeholder
          return `<code class="language-${lang || 'text'}">${code}</code>`;
        }
      });

      const html = marked(content);
      setHtmlContent(html);
    }
  }, [content]);

  if (!content) {
    return <div className="text-gray-500 italic">No content available</div>;
  }

  return (
    <div 
      className="markdown-content max-w-none text-gray-800 leading-relaxed"
      dangerouslySetInnerHTML={{ __html: htmlContent }}
    />
  );
};

export default MarkdownRenderer;