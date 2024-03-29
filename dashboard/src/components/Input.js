import React from 'react'
import { Input } from '@windmill/react-ui';
import classNames from 'classnames'

function BorderedInput({
  className,
  type,
  value,
  placeholder,
  onChange
}) {
  const baseStyle = 'bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500'

  const cls = classNames(baseStyle, className)
  return (
    <Input
    className={cls}
    type={type}
    placeholder={placeholder}
    value={value}
    onChange={onChange}
  />
  )
}

export default BorderedInput;
