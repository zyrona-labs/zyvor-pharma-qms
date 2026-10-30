import { FlaskConical, Package, ShieldCheck, AlertTriangle } from 'lucide-react'
import GlassCard from '../components/GlassCard'

export default function Dashboard() {
  return (
    <div className="max-w-7xl mx-auto px-6 py-10">
      <div className="mb-8">
        <h1 className="text-3xl font-semibold tracking-tight">Overview</h1>
        <p className="text-zinc-400 mt-1">Production batch → laboratory → quality approval → distribution → recall</p>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 mb-10">
        <GlassCard>
          <div className="flex items-center gap-3">
            <div className="p-2 rounded-2xl bg-blue-500/20 text-blue-400"><Package className="w-5 h-5" /></div>
            <div>
              <div className="text-2xl font-semibold">12</div>
              <div className="text-xs text-zinc-400">Active batches</div>
            </div>
          </div>
        </GlassCard>
        <GlassCard>
          <div className="flex items-center gap-3">
            <div className="p-2 rounded-2xl bg-purple-500/20 text-purple-400"><FlaskConical className="w-5 h-5" /></div>
            <div>
              <div className="text-2xl font-semibold">7</div>
              <div className="text-xs text-zinc-400">In laboratory</div>
            </div>
          </div>
        </GlassCard>
        <GlassCard>
          <div className="flex items-center gap-3">
            <div className="p-2 rounded-2xl bg-emerald-500/20 text-emerald-400"><ShieldCheck className="w-5 h-5" /></div>
            <div>
              <div className="text-2xl font-semibold">3</div>
              <div className="text-xs text-zinc-400">Awaiting QA</div>
            </div>
          </div>
        </GlassCard>
        <GlassCard>
          <div className="flex items-center gap-3">
            <div className="p-2 rounded-2xl bg-amber-500/20 text-amber-400"><AlertTriangle className="w-5 h-5" /></div>
            <div>
              <div className="text-2xl font-semibold">1</div>
              <div className="text-xs text-zinc-400">Open deviations</div>
            </div>
          </div>
        </GlassCard>
      </div>

      <div className="glass rounded-3xl p-6">
        <h2 className="font-medium mb-4">Recent batches</h2>
        <div className="space-y-3 text-sm">
          {[
            { num: 'B-2026-0042', product: 'Paracetamol 500 mg', status: 'Laboratory', color: 'text-purple-400' },
            { num: 'B-2026-0041', product: 'Amoxicillin 250 mg', status: 'Quality review', color: 'text-amber-400' },
            { num: 'B-2026-0040', product: 'Ibuprofen 400 mg', status: 'Released', color: 'text-emerald-400' },
          ].map(b => (
            <div key={b.num} className="flex items-center justify-between py-2 border-b border-white/5 last:border-0">
              <div>
                <div className="font-medium">{b.num}</div>
                <div className="text-zinc-400 text-xs">{b.product}</div>
              </div>
              <span className={`text-xs font-medium ${b.color}`}>{b.status}</span>
            </div>
          ))}
        </div>
      </div>
    </div>
  )
}
