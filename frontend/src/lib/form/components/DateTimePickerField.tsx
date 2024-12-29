import { TextFieldProps } from '@mui/material'
import { DateTimePicker, DateTimePickerProps, LocalizationProvider } from '@mui/x-date-pickers'
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs'
import dayjs from 'dayjs'
import 'dayjs/locale/ja'
import { Controller, FieldPath, FieldValues, useFormContext } from 'react-hook-form'

interface Props<TFieldValues extends FieldValues> {
  name: FieldPath<TFieldValues>
  label: string
  views?: ('year' | 'month' | 'day' | 'hours' | 'minutes' | 'seconds')[]
  textFieldProps: Omit<TextFieldProps, 'onChange' | 'value'>
  pickerProps: DateTimePickerProps<dayjs.Dayjs>
}

dayjs.locale('ja')

const DateTimePickerField = <TFieldValues extends FieldValues>({
  name,
  label,
  views,
  textFieldProps,
  pickerProps
}: Props<TFieldValues>) => {
  const { formState, control } = useFormContext()

  return (
    <LocalizationProvider dateAdapter={AdapterDayjs}>
      <Controller
        name={name}
        control={control}
        render={({ field: { onChange, value, ref } }) => (
          <DateTimePicker
            label={label}
            onChange={(newValue) => onChange(newValue?.toISOString() ?? null)}
            value={value ? dayjs(value) : null}
            inputRef={ref}
            views={views}
            slotProps={{
              textField: {
                error: !!formState.errors[name],
                helperText: formState.errors[name]?.message as string,
                ...textFieldProps
              },
              field: { clearable: true, onClear: () => onChange(null) }
            }}
            {...pickerProps}
          />
        )}
      />
    </LocalizationProvider>
  )
}

export default DateTimePickerField
