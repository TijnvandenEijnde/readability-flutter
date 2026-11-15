#
# To learn more about a Podspec see http://guides.cocoapods.org/syntax/podspec.html.
# Run `pod lib lint readability.podspec` to validate before publishing.
#
Pod::Spec.new do |s|
  s.name             = 'readability'
  s.version          = '0.0.1'
  s.summary          = 'A new Flutter FFI plugin project.'
  s.description      = <<-DESC
A new Flutter FFI plugin project.
                       DESC
  s.homepage         = 'http://example.com'
  s.license          = { :file => '../LICENSE' }
  s.author           = { 'Your Company' => 'email@example.com' }

  # The sources and the xcframework live in readability/, which is shared with
  # the Swift Package Manager package (readability/Package.swift).
  s.source           = { :path => '.' }
  s.source_files = 'readability/Sources/readability/**/*.{c,h}'
  s.dependency 'Flutter'
  s.platform = :ios, '13.0'

  s.public_header_files = 'readability/Sources/readability/include/**/*.h'
  s.vendored_frameworks = 'readability/libreadability.xcframework'

  # Flutter.framework does not contain a i386 slice.
  s.pod_target_xcconfig = { 'DEFINES_MODULE' => 'YES', 'EXCLUDED_ARCHS[sdk=iphonesimulator*]' => 'i386' }
end
