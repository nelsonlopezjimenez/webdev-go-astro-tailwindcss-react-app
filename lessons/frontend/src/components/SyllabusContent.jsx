export default function SyllabusContent() {
  return (
    <div className="prose prose-lg max-w-none">
      <div className="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
        
        <section className="mb-8">
          <h2 className="text-2xl font-bold text-gray-900 mb-4">STATEMENT ON ACADEMIC INTEGRITY</h2>
          <p className="text-gray-700 leading-relaxed">
            Edmonds CC students shall demonstrate Academic Integrity. I am expected to report all violations of Academic Integrity (cheating and plagiarism) to the College. The College's database of such incidents will be monitored by the Office of the Vice President for Student Services. Data will be maintained for three years. Evidence of repeat incidents will result in additional action by the Office of the Vice President for Student Services as governed by the Student Code of Conduct. In this class, cheating and plagiarism will result in an assignment or grade penalty ranging from "0" on an assignment to an "F" in the course.
          </p>
        </section>

        <section className="mb-8">
          <h2 className="text-2xl font-bold text-gray-900 mb-4">FINAL EXAM AND LAST MEETING OF CLASS</h2>
          <ul className="list-disc list-inside text-gray-700 space-y-2">
            <li>The final exam will be in-class on <strong>WEEK TENTH, Thursday June 12, 2020</strong>, provided that Covid-19 situation is contained.</li>
          </ul>
        </section>

        <section className="mb-8">
          <h2 className="text-2xl font-bold text-gray-900 mb-4">SERVICES FOR STUDENTS WITH DISABILITIES</h2>
          <p className="text-gray-700 leading-relaxed">
            If you require an accommodation for a disability, please contact the <strong className="text-blue-600">Dean of Corrections Education Brent Arbes</strong>
          </p>
        </section>

        <section className="mb-8">
          <h2 className="text-2xl font-bold text-gray-900 mb-4">COURSE EXPECTATIONS</h2>
          
          <div className="grid md:grid-cols-2 gap-6">
            <div className="bg-blue-50 p-4 rounded-lg">
              <h3 className="text-lg font-semibold text-blue-900 mb-3">Students:</h3>
              <ul className="list-disc list-inside text-gray-700 space-y-2">
                <li>Submit all assignments by the due dates specified in Canvas.</li>
                <li>Send all assignments in Canvas.</li>
                <li>Ensure your name is on your assignment at the top of the first page.</li>
              </ul>
            </div>

            <div className="bg-green-50 p-4 rounded-lg">
              <h3 className="text-lg font-semibold text-green-900 mb-3">Instructor:</h3>
              <ul className="list-disc list-inside text-gray-700 space-y-2">
                <li>I will return email messages within 24 hours, usually much sooner.</li>
                <li>I will grade and return all assignments within seven/ten days.</li>
                <li>I will monitor the discussion board on a daily basis, provided that students have access to it.</li>
                <li>I will post all grades and points on the canvas course student grade book.</li>
              </ul>
            </div>
          </div>
        </section>

        <section className="bg-yellow-50 p-4 rounded-lg border-l-4 border-yellow-400">
          <h3 className="text-lg font-semibold text-yellow-900 mb-2">Important Note</h3>
          <p className="text-gray-700 leading-relaxed">
            Successful completion of student responsibilities in this class requires class attendance and online access to Canvas. From April 13 through May 5 remote mode, after May 5 face-to-face mode. You are expected to attend class and log in to your laptop. Instructions for access to Canvas and student technical support{' '}
            <a 
              href="http://www.edcc.edu/elearning" 
              target="_blank" 
              rel="noopener noreferrer"
              className="text-blue-600 underline hover:text-blue-800"
            >
              (full URL: http://www.edcc.edu/elearning)
            </a>.
          </p>
        </section>
      </div>
    </div>
  );
}