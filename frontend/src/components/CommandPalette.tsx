import { useEffect, useState } from 'react'
import { Search } from 'lucide-react'

export default function CommandPalette() {
  const [open, setOpen] = useState(false)

  useEffect(() => {
    const onKey = (e: KeyboardEvent) => {
      if ((e.metaKey || e.ctrlKey) && e.key === 'k') {
        e.preventDefault()
        setOpen(o => !o)
      }
      if (e.key === 'Escape') setOpen(false)
    }
    window.addEventListener('keydown', onKey)
    return () => window.removeEventListener('keydown', onKey)
  }, [])

  if (!open) return null

  return (
    <div className="fixed inset-0 z-[100] flex items-start justify-center pt-[20vh] bg-black/40 backdrop-blur-sm" onClick={() => setOpen(false)}>
      <div className="w-full max-w-xl glass rounded-3xl overflow-hidden shadow-2xl" onClick={e => e.stopPropagation()}>
        <div className="flex items-center gap-3 px-4 py-3 border-b border-white/10">
          <Search className="w-5 h-5 text-zinc-400" />
          <input
            autoFocus
            placeholder="Search batches, samples, deviations, SOPs…"
            className="flex-1 bg-transparent outline-none text-sm placeholder:text-zinc-500"
          />
        </div>
        <div className="p-2 text-sm text-zinc-400">
          <div className="px-3 py-2 rounded-2xl hover:bg-white/5 cursor-pointer">Open batch B-2026-0042</div>
          <div className="px-3 py-2 rounded-2xl hover:bg-white/5 cursor-pointer">Create deviation</div>
          <div className="px-3 py-2 rounded-2xl hover:bg-white/5 cursor-pointer">View audit trail</div>
        </div>
      </div>
    </div>
  )
}
