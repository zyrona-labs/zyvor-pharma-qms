import { Routes, Route } from 'react-router-dom'
import TopNav from './components/TopNav'
import CommandPalette from './components/CommandPalette'
import Dashboard from './pages/Dashboard'
import Batches from './pages/Batches'

export default function App() {
  return (
    <div className="min-h-screen">
      <TopNav />
      <CommandPalette />
      <main className="pt-16">
        <Routes>
          <Route path="/" element={<Dashboard />} />
          <Route path="/batches" element={<Batches />} />
        </Routes>
      </main>
    </div>
  )
}
