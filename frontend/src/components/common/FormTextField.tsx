import { TextField, TextFieldProps } from '@mui/material'
import { FieldPath, FieldValues, UseFormReturn } from 'react-hook-form'

interface Props<TFieldValues extends FieldValues>
  extends Omit<TextFieldProps, 'onChange' | 'value'> {
  name: FieldPath<TFieldValues>
  form: UseFormReturn<TFieldValues>
}

const FormTextField = <TFieldValues extends FieldValues>({
  name,
  form,
  ...textFieldProps
}: Props<TFieldValues>) => {
  return (
    <TextField
      {...form.register(name)}
      {...textFieldProps}
      error={!!form.formState.errors[name]}
      helperText={form.formState.errors[name]?.message as string}
      slotProps={{ inputLabel: { shrink: !!form.watch(name) } }}
    />
  )
}

export default FormTextField
