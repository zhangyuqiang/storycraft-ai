export type Project = {
  id: number
  title: string
  genre: string
  status: string
  description?: string
  summary?: string
  chapters?: number
}

export type ProjectCreatePayload = {
  title: string
  description?: string
  genre?: string
  visibility?: string
}

export type ProjectUpdatePayload = {
  title?: string
  description?: string
  genre?: string
  visibility?: string
  status?: string
}

export async function fetchProjects(): Promise<Project[]> {
  const res = await fetch('http://localhost:8080/api/v1/projects', {
    method: 'GET',
    headers: { 'Content-Type': 'application/json' },
  })

  if (!res.ok) {
    throw new Error(`Failed to fetch projects: ${res.status}`)
  }

  const json = await res.json()
  const items = Array.isArray(json.items) ? json.items : []

  return items.map((item: any) => ({
    id: Number(item.id ?? 0),
    title: String(item.title ?? 'Untitled project'),
    genre: String(item.genre ?? 'unknown'),
    status: String(item.status ?? 'draft'),
    description: String(item.description ?? ''),
    summary: String(item.description ?? ''),
    chapters: Number(item.chapters ?? 0),
  }))
}

export async function createProject(payload: ProjectCreatePayload) {
  const res = await fetch('http://localhost:8080/api/v1/projects', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload),
  })

  if (!res.ok) {
    throw new Error(`Failed to create project: ${res.status}`)
  }

  return res.json()
}

export async function updateProject(id: number, payload: ProjectUpdatePayload) {
  const res = await fetch(`http://localhost:8080/api/v1/projects/${id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload),
  })

  if (!res.ok) {
    throw new Error(`Failed to update project: ${res.status}`)
  }

  return res.json()
}

export async function generateStory(payload: { topic: string; genre: string; tone?: string }) {
  const res = await fetch('http://localhost:8080/api/v1/ai/story/generate', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload),
  })

  if (!res.ok) {
    throw new Error(`Failed to generate story: ${res.status}`)
  }

  return res.json()
}

export async function generateShot(payload: { prompt: string; style?: string }) {
  const res = await fetch('http://localhost:8080/api/v1/ai/shot/generate', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload),
  })

  if (!res.ok) {
    throw new Error(`Failed to generate shot: ${res.status}`)
  }

  return res.json()
}
