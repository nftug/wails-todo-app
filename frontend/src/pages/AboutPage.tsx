import { Box, Typography } from '@mui/material'

const AboutPage: React.FC = () => {
  return (
    <Box sx={{ height: 'calc(100vh - 64px - 16px)' }}>
      <Box display="flex" justifyContent="center" alignItems="center" height={1}>
        <Typography variant="h2" color="textSecondary">
          Todo App
        </Typography>
      </Box>
    </Box>
  )
}

export default AboutPage
