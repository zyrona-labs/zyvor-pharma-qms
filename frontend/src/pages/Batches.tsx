export default function Batches() {
  return (
    <div className="max-w-7xl mx-auto px-6 py-10">
      <div className="flex items-center justify-between mb-8">
        <div>
          <h1 className="text-3xl font-semibold tracking-tight">Batches</h1>
          <p className="text-zinc-400 mt-1">Production records and genealogy</p>
        </div>
        <button className="px-4 py-2 rounded-2xl bg-white text-black text-sm font-medium hover:bg-zinc-200 transition">
          New batch
        </button>
      </div>

      <div className="glass rounded-3xl overflow-hidden">
        <table className="w-full text-sm">
          <thead className="text-left text-zinc-400 border-b border-white/10">
            <tr>
              <th className="px-5 py-3 font-medium">Number</th>
              <th className="px-5 py-3 font-medium">Product</th>
              <th className="px-5 py-3 font-medium">Status</th>
              <th className="px-5 py-3 font-medium">Yield</th>
            </tr>
          </thead>
          <tbody>
            <tr className="border-b border-white/5 hover:bg-white/5">
              <td className="px-5 py-4 font-medium">B-2026-0042</td>
              <td className="px-5 py-4">Paracetamol 500 mg</td>
              <td className="px-5 py-4 text-purple-400">Laboratory</td>
              <td className="px-5 py-4">—</td>
            </tr>
            <tr className="border-b border-white/5 hover:bg-white/5">
              <td className="px-5 py-4 font-medium">B-2026-0041</td>
              <td className="px-5 py-4">Amoxicillin 250 mg</td>
              <td className="px-5 py-4 text-amber-400">Quality review</td>
              <td className="px-5 py-4">98.2%</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  )
}
