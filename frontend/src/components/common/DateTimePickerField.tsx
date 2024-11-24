import { TextFieldProps } from '@mui/material'
import { DateTimePicker, LocalizationProvider } from '@mui/x-date-pickers'
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs'
import dayjs from 'dayjs'
import { Controller, FieldPath, FieldValues, UseFormReturn } from 'react-hook-form'

interface Props<TFieldValues extends FieldValues>
  extends Omit<TextFieldProps, 'onChange' | 'value'> {
  name: FieldPath<TFieldValues>
  form: UseFormReturn<TFieldValues>
  label: string
  views?: ('year' | 'month' | 'day' | 'hours' | 'minutes' | 'seconds')[]
}

const DateTimePickerField = <TFieldValues extends FieldValues>({
  name,
  form,
  label,
  views,
  ...textFieldProps
}: Props<TFieldValues>) => {
  return (
    <LocalizationProvider dateAdapter={AdapterDayjs}>
      <Controller
        name={name}
        control={form.control}
        render={({ field: { onChange, value, ref } }) => (
          <DateTimePicker
            label={label}
            onChange={onChange}
            value={value ? dayjs(value) : null}
            inputRef={ref}
            views={views}
            slotProps={{
              textField: {
                error: !!form.formState.errors[name],
                helperText: form.formState.errors[name]?.message as string,
                ...textFieldProps
              }
            }}
          />
        )}
      />
    </LocalizationProvider>
  )
}

export default DateTimePickerField
