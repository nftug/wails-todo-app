import { createContext, useState } from 'react'

export const DrawerContext = createContext<boolean>(undefined!)
export const DrawerDispatchContext = createContext<React.Dispatch<React.SetStateAction<boolean>>>(
  undefined!
)

export const HeaderProvider = ({ children }: { children?: React.ReactNode }) => {
  const [drawerOpened, setDrawerOpened] = useState(false)

  return (
    <DrawerContext.Provider value={drawerOpened}>
      <DrawerDispatchContext.Provider value={setDrawerOpened}>
        {children}
      </DrawerDispatchContext.Provider>
    </DrawerContext.Provider>
  )
}
