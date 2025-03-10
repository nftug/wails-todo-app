import { enums, todo } from '@wailsjs/go/models'
import * as yup from 'yup'

export const todoFieldSchema: yup.ObjectSchema<todo.CreateCommand> = yup.object().shape({
  title: yup.string().required('タイトルは必須です').default(''),
  description: yup.string().optional().default(''),
  dueDate: yup.string().optional().nullable(),
  initialStatus: yup.mixed<enums.StatusValue>().optional()
})

export type TodoFieldSchemaType = yup.InferType<typeof todoFieldSchema>
