import {
  Container,
  FormControl,
  FormControlLabel,
  FormLabel,
  Radio,
  RadioGroup,
  Stack,
  Typography,
  useColorScheme
} from '@mui/material'

const SettingsPage: React.FC = () => {
  const { mode, setMode } = useColorScheme()
  if (!mode) {
    return null
  }

  return (
    <Container sx={{ marginTop: 5 }}>
      <Stack spacing={3}>
        <Typography variant="h4">設定</Typography>

        <FormControl>
          <FormLabel>テーマ</FormLabel>
          <RadioGroup
            row
            value={mode}
            onChange={(event) => setMode(event.target.value as 'system' | 'light' | 'dark')}
          >
            <FormControlLabel value="system" control={<Radio />} label="システムに合わせる" />
            <FormControlLabel value="light" control={<Radio />} label="ライトモード" />
            <FormControlLabel value="dark" control={<Radio />} label="ダークモード" />
          </RadioGroup>
        </FormControl>
      </Stack>
    </Container>
  )
}

export default SettingsPage
