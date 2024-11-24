import { TextFieldProps } from '@mui/material'
import { DateTimePicker, LocalizationProvider } from '@mui/x-date-pickers'
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs'
import dayjs from 'dayjs'
import { Control, Controller, FieldPath, FieldValues } from 'react-hook-form'

interface DateTimePickerFieldProps<TFieldValues extends FieldValues>
  extends Omit<TextFieldProps, 'onChange' | 'value'> {
  name: FieldPath<TFieldValues>
  control: Control<TFieldValues>
  label: string
  views?: ('year' | 'month' | 'day' | 'hours' | 'minutes' | 'seconds')[]
}

const DateTimePickerField = <TFieldValues extends FieldValues>({
  name,
  control,
  label,
  views,
  ...textFieldProps
}: DateTimePickerFieldProps<TFieldValues>) => {
  return (
    <LocalizationProvider dateAdapter={AdapterDayjs}>
      <Controller
        name={name}
        control={control}
        render={({ field: { onChange, value, ref } }) => (
          <DateTimePicker
            label={label}
            onChange={onChange}
            value={value ? dayjs(value) : null}
            inputRef={ref}
            views={views}
            slotProps={{
              textField: textFieldProps
            }}
          />
        )}
      />
    </LocalizationProvider>
  )
}

export default DateTimePickerField
