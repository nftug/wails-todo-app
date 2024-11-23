import { enums } from '../types/wailsjs/go/models'

interface ApiErrorData {
  field: string
  message: string
}

interface ApiErrorJson {
  code: enums.ErrorCode
  data?: ApiErrorData
}

export class ApiError extends Error {
  private constructor(
    public readonly code: enums.ErrorCode,
    public readonly data?: ApiErrorData
  ) {
    super(code)
  }

  static create(errJson: ApiErrorJson) {
    return new ApiError(errJson.code, errJson.data)
  }
}

export const handleApiError = async <T>(callback: () => Promise<T>) => {
  try {
    return await callback()
  } catch (e) {
    if (typeof e === 'string') {
      throw ApiError.create(JSON.parse(e) as ApiErrorJson)
    } else {
      throw e
    }
  }
}
