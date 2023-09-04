import { create } from 'zustand'

type EditedTask = {
  id: number
  title: string
}

type State = {
  // ステート
  editedTask: EditedTask
  // 関数
  updateEditedTask: (payload: EditedTask) => void
  resetEditedTask: () => void
}

const useStore = create<State>((set) => ({
  editedTask: { id: 0, title: '' },
  updateEditedTask: (payload) => set({ editedTask: payload }),
  resetEditedTask: () => set({ editedTask: { id: 0, title: '' } }),
}))

export default useStore
