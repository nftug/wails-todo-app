import { Theme } from '@emotion/react'
import { SxProps } from '@mui/material'

export const overflowEllipsisStyle: SxProps<Theme> = {
  overflow: 'hidden',
  textOverflow: 'ellipsis',
  whiteSpace: 'nowrap',
  maxWidth: 1
} as const
