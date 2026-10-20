export type Project = {
  id: number
  title: string
  genre: string
  status: string
}

export async function fetchProjects(): Promise<Project[]> {
  return [
    { id: 1, title: 'Moonfall City', genre: 'Fantasy', status: 'In progress' },
    { id: 2, title: 'Glass Horizon', genre: 'Sci-fi', status: 'Draft' },
  ]
}
