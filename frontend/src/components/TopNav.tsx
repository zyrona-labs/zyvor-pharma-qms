import { Search, Bell, User } from 'lucide-react'
import { useState } from 'react'

export default function TopNav() {
  const [env, setEnv] = useState<'test' | 'live'>('test')

  return (
    <header className="fixed top-0 inset-x-0 z-50 h-14 glass flex items-center justify-between px-6">
      <div className="flex items-center gap-8">
        <div className="flex items-center gap-2">
          <div className="w-8 h-8 rounded-2xl bg-gradient-to-br from-orange-500 to-amber-600 flex items-center justify-center text-white font-semibold text-sm">
            Z
          </div>
          <span className="font-semibold tracking-tight">Zyvor Pharma QMS</span>
        </div>
        <nav className="hidden md:flex items-center gap-6 text-sm text-zinc-300">
          <a href="/" className="hover:text-white transition">Overview</a>
          <a href="/batches" className="hover:text-white transition">Batches</a>
          <a href="#" className="hover:text-white transition">Laboratory</a>
          <a href="#" className="hover:text-white transition">Quality</a>
          <a href="#" className="hover:text-white transition">Traceability</a>
        </nav>
      </div>

      <div className="flex items-center gap-3">
        <button className="flex items-center gap-2 px-3 py-1.5 rounded-2xl glass-light text-sm text-zinc-300 hover:bg-white/10 transition">
          <Search className="w-4 h-4" />
          <span className="hidden sm:inline">Search</span>
          <kbd className="hidden sm:inline text-[10px] px-1.5 py-0.5 rounded bg-white/10">⌘K</kbd>
        </button>

        <div className="flex items-center gap-1 text-xs font-medium">
          <button
            onClick={() => setEnv('test')}
            className={`px-2.5 py-1 rounded-xl transition ${env === 'test' ? 'bg-amber-500/20 text-amber-300' : 'text-zinc-500 hover:text-zinc-300'}`}
          >
            Test
          </button>
          <button
            onClick={() => setEnv('live')}
            className={`px-2.5 py-1 rounded-xl transition ${env === 'live' ? 'bg-emerald-500/20 text-emerald-300' : 'text-zinc-500 hover:text-zinc-300'}`}
          >
            Live
          </button>
        </div>

        <button className="p-2 rounded-2xl glass-light hover:bg-white/10 transition">
          <Bell className="w-4 h-4" />
        </button>
        <button className="p-2 rounded-2xl glass-light hover:bg-white/10 transition">
          <User className="w-4 h-4" />
        </button>
      </div>
    </header>
  )
}
