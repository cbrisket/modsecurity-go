package seclang

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/timtadh/lexmachine/machines"
)

const TkStart = 128
const (
	freetextRegex   = `(([^"\n\\]|(\\.)|(\\\n))+)`
	charactersRegex = `([^, \t\"\n\r]+)`
)

const (
	TkActionAccuracy = iota + TkStart
	TkActionAllow
	TkActionAppend
	TkActionAuditLog
	TkActionBlock
	TkActionCapture
	TkActionChain
	TkActionCtlAuditEngine
	TkActionCtlAuditLogParts
	TkActionCtlBdyJson
	TkActionCtlBdyXml
	TkActionCtlBdyUrlencoded
	TkActionCtlForceReqBodyVar
	TkActionCtlRequestBodyAccess
	TkActionCtlRuleEngine
	TkActionCtlRuleRemoveByTag
	TkActionCtlRuleRemoveById
	TkActionCtlRuleRemoveTargetById
	TkActionCtlRuleRemoveTargetByTag
	TkActionDeny
	TkActionDeprecateVar
	TkActionDrop
	TkActionExec
	TkActionExpireVar
	TkActionId
	TkActionInitcol
	TkActionLogData
	TkActionLog
	TkActionMaturity
	TkActionMsg
	TkActionMultiMatch
	TkActionNoAuditLog
	TkActionNoLog
	TkActionPass
	TkActionPause
	TkActionPhase
	TkActionPrepend
	TkActionProxy
	TkActionRedirect
	TkActionRev
	TkActionSanitiseArg
	TkActionSanitiseMatchedBytes
	TkActionSanitiseMatched
	TkActionSanitiseRequestHeader
	TkActionSanitiseResponseHeader
	TkActionSetenv
	TkActionSetrsc
	TkActionSetsid
	TkActionSetuid
	TkActionSetvar
	TkActionSeverity
	TkActionSeverityValue
	TkActionSkipAfter
	TkActionSkip
	TkActionStatus
	TkActionTag
	TkActionVer
	TkActionXmlns
	TkActionTransformationBase64Encode
	TkActionTransformationBase64Decode
	TkActionTransformationBase64DecodeExt
	TkActionTransformationCmdLine
	TkActionTransformationCompressWhitespace
	TkActionTransformationEscapeSeqDecode
	TkActionTransformationCssDecode
	TkActionTransformationHexEncode
	TkActionTransformationHexDecode
	TkActionTransformationHtmlEntityDecode
	TkActionTransformationJsDecode
	TkActionTransformationLength
	TkActionTransformationLowercase
	TkActionTransformationMd5
	TkActionTransformationNone
	TkActionTransformationNormalisePath
	TkActionTransformationNormalisePathWin
	TkActionTransformationParityEven7Bit
	TkActionTransformationParityOdd7Bit
	TkActionTransformationParityZero7Bit
	TkActionTransformationRemoveComments
	TkActionTransformationRemoveCommentsChar
	TkActionTransformationRemoveNulls
	TkActionTransformationRemoveWhitespace
	TkActionTransformationReplaceComments
	TkActionTransformationReplaceNulls
	TkActionTransformationSha1
	TkActionTransformationSqlHexDecode
	TkActionTransformationTrim
	TkActionTransformationTrimLeft
	TkActionTransformationTrimRight
	TkActionTransformationUppercase
	TkActionTransformationUrlEncode
	TkActionTransformationUrlDecode
	TkActionTransformationUrlDecodeUni
	TkActionTransformationUtf8ToUnicode
	TkVariableArgsCombinedSize
	TkVariableArgsGetNames
	TkVariableArgsNames
	TkVariableArgsPostNames
	TkVariableAuthType
	TkVariableFilesCombinedSize
	TkVariableFilesTmpNames
	TkVariableFullRequest
	TkVariableFullRequestLength
	TkVariableGlobal
	TkVariableInboundDataError
	TkVariableMatchedVar
	TkVariableMatchedVarName
	TkVariableMultipartBoundaryQuoted
	TkVariableMultipartBoundaryWhitespace
	TkVariableMultipartCrlfLfLines
	TkVariableMultipartDataAfter
	TkVariableMultipartDataBefore
	TkVariableMultipartFileLimitExceeded
	TkVariableMultipartFilename
	TkVariableMultipartHeaderFolding
	TkVariableMultipartInvalidHeaderFolding
	TkVariableMultipartInvalidPart
	TkVariableMultipartInvalidQuoting
	TkVariableMultipartLfLine
	TkVariableMultipartMissingSemicolon
	TkVariableMultipartSemicolonMissing
	TkVariableMultipartName
	TkVariableMultipartStrictError
	TkVariableMultipartUnmatchedBoundary
	TkVariableOutboundDataError
	TkVariablePathInfo
	TkVariableQueryString
	TkVariableRemoteAddr
	TkVariableRemoteHost
	TkVariableRemotePort
	TkVariableReqbodyError
	TkVariableReqbodyErrorMsg
	TkVariableReqbodyProcessorError
	TkVariableReqbodyProcessorErrorMsg
	TkVariableReqbodyProcessor
	TkVariableRequestBasename
	TkVariableRequestBody
	TkVariableRequestBodyLength
	TkVariableRequestFileName
	TkVariableRequestHeadersNames
	TkVariableRequestLine
	TkVariableRequestMethod
	TkVariableRequestProtocol
	TkVariableRequestUri
	TkVariableRequestUriRaw
	TkVariableResource
	TkVariableResponseBody
	TkVariableResponseContentLength
	TkVariableResponseContentType
	TkVariableResponseHeadersNames
	TkVariableResponseProtocol
	TkVariableResponseStatus
	TkVariableServerAddr
	TkVariableServerName
	TkVariableServerPort
	TkVariableSessionId
	TkVariableUniqueId
	TkVariableUrlEncodedError
	TkVariableUserId
	TkVariableWebserverErrorLog
	TkVariableArgs
	TkVariableArgsPost
	TkVariableArgsGet
	TkVariableFilesSizes
	TkVariableFilesNames
	TkVariableFilesTmpContent
	TkVariableMatchedVarsNames
	TkVariableMatchedVars
	TkVariableFiles
	TkVariableRequestCookies
	TkVariableRequestHeaders
	TkVariableResponseHeaders
	TkVariableGeo
	TkVariableRequestCookiesNames
	TkVariableRule
	TkVariableSession
	TkVariableIp
	TkVariableUser
	TkVariableStatus
	TkVariableStatusLine
	TkVariableTx
	TkVariableWebAppId
	TkRunTimeVarBld
	TkRunTimeVarDur
	TkRunTimeVarEnv
	TkRunTimeVarHsv
	TkRunTimeVarRemoteUser
	TkRunTimeVarTime
	TkRunTimeVarTimeDay
	TkRunTimeVarTimeEpoch
	TkRunTimeVarTimeHour
	TkRunTimeVarTimeMin
	TkRunTimeVarTimeMon
	TkRunTimeVarTimeSec
	TkRunTimeVarTimeWday
	TkRunTimeVarTimeYear
	TkRunTimeVarXml
	TkVarExclusion
	TkVarCount
	TkOperatorBeginsWith
	TkOperatorContains
	TkOperatorContainsWord
	TkOperatorDetectSqli
	TkOperatorDetectXss
	TkOperatorEndsWith
	TkOperatorEq
	TkOperatorFuzzyHash
	TkOperatorGe
	TkOperatorGeolookup
	TkOperatorGsbLookup
	TkOperatorGt
	TkOperatorInspectFile
	TkOperatorIpMatchFromFile
	TkOperatorIpMatch
	TkOperatorLe
	TkOperatorLt
	TkOperatorPmFromFile
	TkOperatorPm
	TkOperatorRbl
	TkOperatorRsub
	TkOperatorRx
	TkOperatorStrEq
	TkOperatorStrMatch
	TkOperatorUnconditionalMatch
	TkOperatorValidateByteRange
	TkOperatorValidateDtd
	TkOperatorValidateHash
	TkOperatorValidateSchema
	TkOperatorValidateUrlEncoding
	TkOperatorValidateUtf8Encoding
	TkOperatorVerifyCc
	TkOperatorVerifyCpf
	TkOperatorVerifySsn
	TkOperatorVerifySvnr
	TkOperatorWithin
	TkAuditParts
	TkColFreeTextSpaceComma
	TkColName
	TkConfigComponentSig
	TkConfigSecServerSig
	TkConfigSecWebAppId
	TkConfigSecCacheTransformations
	TkConfigSecChrootDir
	TkConfigConnEngine
	TkConfigSecHashEngine
	TkConfigSecHashKey
	TkConfigSecHashParam
	TkConfigSecHashMethodRx
	TkConfigSecHashMethodPm
	TkConfigContentInjection
	TkConfigSecArgumentSeparator
	TkConfigDirAuditDir
	TkConfigDirAuditDirMod
	TkConfigDirAuditEng
	TkConfigDirAuditFleMod
	TkConfigDirAuditLog2
	TkConfigDirAuditLog
	TkConfigDirAuditLogFmt
	TkConfigDirAuditLogP
	TkConfigDirAuditSts
	TkConfigDirAuditTpe
	TkConfigDirDebugLog
	TkConfigDirDebugLvl
	TkConfigDirGeoDb
	TkConfigDirGsbDb
	TkConfigSecGuardianLog
	TkConfigSecInterceptOnError
	TkConfigSecConnRStateLimit
	TkConfigSecConnWStateLimit
	TkConfigSecSensorId
	TkConfigSecRuleInheritance
	TkConfigSecRulePerfTime
	TkConfigSecStreamInBodyInspection
	TkConfigSecStreamOutBodyInspection
	TkConfigDirPcreMatchLimit
	TkConfigDirPcreMatchLimitRecursion
	TkConfigDirReqBody
	TkConfigDirReqBodyInMemoryLimit
	TkConfigDirReqBodyLimit
	TkConfigDirReqBodyLimitAction
	TkConfigDirReqBodyNoFilesLimit
	TkConfigDirResBody
	TkConfigDirResBodyLimit
	TkConfigDirResBodyLimitAction
	TkConfigDirRuleEng
	TkConfigDirSecAction
	TkConfigDirSecDefaultAction
	TkConfigSecDisableBackendCompress
	TkConfigDirSecMarker
	TkConfigDirUnicodeMapFile
	TkConfigInclude
	TkConfigSecCollectionTimeout
	TkConfigSecHttpBlkey
	TkConfigSecRemoteRules
	TkConfigSecRemoteRulesFailAction
	TkConfigSecRemoveRulesById
	TkConfigSecRemoveRulesByMsg
	TkConfigSecRemoveRulesByTag
	TkConfigSecUpdateTargetByTag
	TkConfigSecUpdateTargetByMsg
	TkConfigSecUpdateTargetById
	TkConfigSecUpdateActionById
	TkConfigUpdloadKeepFiles
	TkConfigUpdloadSaveTmpFiles
	TkConfigUploadDir
	TkConfigUploadFileLimit
	TkConfigUploadFileMode
	TkConfigValueAbort
	TkConfigValueDetc
	TkConfigValueHttps
	TkConfigValueNumber
	TkConfigValueOff
	TkConfigValueOn
	TkConfigValueParallel
	TkConfigValuePath
	TkConfigValueProcessPartial
	TkConfigValueReject
	TkConfigValueRelevantOnly
	TkConfigValueSerial
	TkConfigValueWarn
	TkConfigXmlExternalEntity
	TkConfigDirResponseBodyMp
	TkConfigDirResponseBodyMpClear
	TkConfigDirSecArgSep
	TkConfigDirSecCookieFormat
	TkConfigSecCookiev0Separator
	TkConfigDirSecDataDir
	TkConfigDirSecStatusEngine
	TkConfigDirSecTmpDir
	TkDictElement
	TkDictElementWithPipe
	TkDictElementNoPipe
	TkDictElementNoMacro
	TkDictElementTwo
	TkDictElementTwoQuoted
	TkDictElementTwo2
	TkDirective
	TkDirectiveSecrulescript
	TkQuoteButScaped
	TkDoubleQuoteButScaped
	TkCommaButScaped
	TkStartMacroVariable
	TkNot
	TkRemoveRuleBy
	TkVarFreeTextQuote
	TkVarFreeTextSpace
	TkVarFreeTextSpaceComma
	TkJson
	TkNative
	TkNewLine
	TkEquals
	TkEqualsPlus
	TkEqualsMinus
	TkQuotationMark
	TkComma
	TkFreeText
	TkCharacters
	TkArgument
)

var TkRegexMap = map[int]string{
	TkActionAccuracy:                         `(accuracy)`,
	TkActionAllow:                            `((allow:(REQUEST|PHASE))|(phase:'(REQUEST|PHASE)')|(allow))`,
	TkActionAppend:                           `(append)`,
	TkActionAuditLog:                         `(auditlog)`,
	TkActionBlock:                            `(block)`,
	TkActionCapture:                          `(capture)`,
	TkActionChain:                            `(chain)`,
	TkActionCtlAuditEngine:                   `(ctl:auditEngine)`,
	TkActionCtlAuditLogParts:                 `(ctl:auditLogParts)`,
	TkActionCtlBdyJson:                       `(ctl:requestBodyProcessor=JSON)`,
	TkActionCtlBdyXml:                        `(ctl:requestBodyProcessor=XML)`,
	TkActionCtlBdyUrlencoded:                 `(ctl:requestBodyProcessor=URLENCODED)`,
	TkActionCtlForceReqBodyVar:               `(ctl:forceRequestBodyVariable)`,
	TkActionCtlRequestBodyAccess:             `(ctl:requestBodyAccess)`,
	TkActionCtlRuleEngine:                    `(ctl:ruleEngine)`,
	TkActionCtlRuleRemoveByTag:               `(ctl:ruleRemoveByTag)`,
	TkActionCtlRuleRemoveById:                `(ctl:ruleRemoveById)`,
	TkActionCtlRuleRemoveTargetById:          `(ctl:ruleRemoveTargetById)`,
	TkActionCtlRuleRemoveTargetByTag:         `(ctl:ruleRemoveTargetByTag)`,
	TkActionDeny:                             `(deny)`,
	TkActionDeprecateVar:                     `(deprecatevar)`,
	TkActionDrop:                             `(drop)`,
	TkActionExec:                             `(exec)`,
	TkActionExpireVar:                        `(expirevar)`,
	TkActionId:                               `(id:[0-9]+|id:'[0-9]+')`,
	TkActionInitcol:                          `(initcol)`,
	TkActionLogData:                          `(logdata)`,
	TkActionLog:                              `(log)`,
	TkActionMaturity:                         `(maturity)`,
	TkActionMsg:                              `(msg)`,
	TkActionMultiMatch:                       `(multiMatch)`,
	TkActionNoAuditLog:                       `(noauditlog)`,
	TkActionNoLog:                            `(nolog)`,
	TkActionPass:                             `(pass)`,
	TkActionPause:                            `(pause)`,
	TkActionPhase:                            `((phase:(REQUEST|RESPONSE|LOGGING|[0-9]+))|(phase:'(REQUEST|RESPONSE|LOGGING|[0-9]+)'))`,
	TkActionPrepend:                          `(prepend)`,
	TkActionProxy:                            `(proxy)`,
	TkActionRedirect:                         `(redirect)`,
	TkActionRev:                              `(rev)`,
	TkActionSanitiseArg:                      `(sanitiseArg)`,
	TkActionSanitiseMatchedBytes:             `(sanitiseMatchedBytes)`,
	TkActionSanitiseMatched:                  `(sanitiseMatched)`,
	TkActionSanitiseRequestHeader:            `(sanitiseRequestHeader)`,
	TkActionSanitiseResponseHeader:           `(sanitiseResponseHeader)`,
	TkActionSetenv:                           `(setenv)`,
	TkActionSetrsc:                           `(setrsc)`,
	TkActionSetsid:                           `(setsid)`,
	TkActionSetuid:                           `(setuid)`,
	TkActionSetvar:                           `(setvar)`,
	TkActionSeverity:                         `(severity)`,
	TkActionSeverityValue:                    `((EMERGENCY|ALERT|CRITICAL|ERROR|WARNING|NOTICE|INFO|DEBUG)|[0-9]+)`,
	TkActionSkipAfter:                        `(skipAfter)`,
	TkActionSkip:                             `(skip)`,
	TkActionStatus:                           `(status:[0-9]+)`,
	TkActionTag:                              `(tag)`,
	TkActionVer:                              `(ver)`,
	TkActionXmlns:                            `(xmlns)`,
	TkActionTransformationBase64Encode:       `(t:base64Encode)`,
	TkActionTransformationBase64Decode:       `(t:base64Decode)`,
	TkActionTransformationBase64DecodeExt:    `(t:base64DecodeExt)`,
	TkActionTransformationCmdLine:            `(t:cmdLine)`,
	TkActionTransformationCompressWhitespace: `(t:compressWhitespace)`,
	TkActionTransformationEscapeSeqDecode:    `(t:escapeSeqDecode)`,
	TkActionTransformationCssDecode:          `(t:cssDecode)`,
	TkActionTransformationHexEncode:          `(t:hexEncode)`,
	TkActionTransformationHexDecode:          `(t:hexDecode)`,
	TkActionTransformationHtmlEntityDecode:   `(t:htmlEntityDecode)`,
	TkActionTransformationJsDecode:           `(t:jsDecode)`,
	TkActionTransformationLength:             `(t:length)`,
	TkActionTransformationLowercase:          `(t:lowercase)`,
	TkActionTransformationMd5:                `(t:md5)`,
	TkActionTransformationNone:               `(t:none)`,
	TkActionTransformationNormalisePath:      `(t:(normalisePath|normalizePath))`,
	TkActionTransformationNormalisePathWin:   `(t:(normalisePathWin|normalizePathWin))`,
	TkActionTransformationParityEven7Bit:     `(t:parityEven7bit)`,
	TkActionTransformationParityOdd7Bit:      `(t:parityOdd7bit)`,
	TkActionTransformationParityZero7Bit:     `(t:parityZero7bit)`,
	TkActionTransformationRemoveComments:     `(t:removeComments)`,
	TkActionTransformationRemoveCommentsChar: `(t:removeCommentsChar)`,
	TkActionTransformationRemoveNulls:        `(t:removeNulls)`,
	TkActionTransformationRemoveWhitespace:   `(t:removeWhitespace)`,
	TkActionTransformationReplaceComments:    `(t:replaceComments)`,
	TkActionTransformationReplaceNulls:       `(t:replaceNulls)`,
	TkActionTransformationSha1:               `(t:sha1)`,
	TkActionTransformationSqlHexDecode:       `(t:sqlHexDecode)`,
	TkActionTransformationTrim:               `(t:trim)`,
	TkActionTransformationTrimLeft:           `(t:trimLeft)`,
	TkActionTransformationTrimRight:          `(t:trimRight)`,
	TkActionTransformationUppercase:          `(t:uppercase)`,
	TkActionTransformationUrlEncode:          `(t:urlEncode)`,
	TkActionTransformationUrlDecode:          `(t:urlDecode)`,
	TkActionTransformationUrlDecodeUni:       `(t:urlDecodeUni)`,
	TkActionTransformationUtf8ToUnicode:      `(t:utf8toUnicode)`,
	TkVariableArgsCombinedSize:               `(ARGS_COMBINED_SIZE)`,
	TkVariableArgsGetNames:                   `(ARGS_GET_NAMES)`,
	TkVariableArgsNames:                      `(ARGS_NAMES)`,
	TkVariableArgsPostNames:                  `(ARGS_POST_NAMES)`,
	TkVariableAuthType:                       `(AUTH_TYPE)`,
	TkVariableFilesCombinedSize:              `(FILES_COMBINED_SIZE)`,
	TkVariableFilesTmpNames:                  `(FILES_TMPNAMES)`,
	TkVariableFullRequest:                    `(FULL_REQUEST)`,
	TkVariableFullRequestLength:              `(FULL_REQUEST_LENGTH)`,
	TkVariableGlobal:                         `(GLOBAL)`,
	TkVariableInboundDataError:               `(INBOUND_DATA_ERROR)`,
	TkVariableMatchedVar:                     `(MATCHED_VAR)`,
	TkVariableMatchedVarName:                 `(MATCHED_VAR_NAME)`,
	TkVariableMultipartBoundaryQuoted:        `(MULTIPART_BOUNDARY_QUOTED)`,
	TkVariableMultipartBoundaryWhitespace:    `(MULTIPART_BOUNDARY_WHITESPACE)`,
	TkVariableMultipartCrlfLfLines:           `(MULTIPART_CRLF_LF_LINES)`,
	TkVariableMultipartDataAfter:             `(MULTIPART_DATA_AFTER)`,
	TkVariableMultipartDataBefore:            `(MULTIPART_DATA_BEFORE)`,
	TkVariableMultipartFileLimitExceeded:     `(MULTIPART_FILE_LIMIT_EXCEEDED)`,
	TkVariableMultipartFilename:              `(MULTIPART_FILENAME)`,
	TkVariableMultipartHeaderFolding:         `(MULTIPART_HEADER_FOLDING)`,
	TkVariableMultipartInvalidHeaderFolding:  `(MULTIPART_INVALID_HEADER_FOLDING)`,
	TkVariableMultipartInvalidPart:           `(MULTIPART_INVALID_PART)`,
	TkVariableMultipartInvalidQuoting:        `(MULTIPART_INVALID_QUOTING)`,
	TkVariableMultipartLfLine:                `(MULTIPART_LF_LINE)`,
	TkVariableMultipartMissingSemicolon:      `(MULTIPART_MISSING_SEMICOLON)`,
	TkVariableMultipartSemicolonMissing:      `(MULTIPART_SEMICOLON_MISSING)`,
	TkVariableMultipartName:                  `(MULTIPART_NAME)`,
	TkVariableMultipartStrictError:           `(MULTIPART_STRICT_ERROR)`,
	TkVariableMultipartUnmatchedBoundary:     `(MULTIPART_UNMATCHED_BOUNDARY)`,
	TkVariableOutboundDataError:              `(OUTBOUND_DATA_ERROR)`,
	TkVariablePathInfo:                       `(PATH_INFO)`,
	TkVariableQueryString:                    `(QUERY_STRING)`,
	TkVariableRemoteAddr:                     `(REMOTE_ADDR)`,
	TkVariableRemoteHost:                     `(REMOTE_HOST)`,
	TkVariableRemotePort:                     `(REMOTE_PORT)`,
	TkVariableReqbodyError:                   `(REQBODY_ERROR)`,
	TkVariableReqbodyErrorMsg:                `(REQBODY_ERROR_MSG)`,
	TkVariableReqbodyProcessorError:          `(REQBODY_PROCESSOR_ERROR)`,
	TkVariableReqbodyProcessorErrorMsg:       `(REQBODY_PROCESSOR_ERROR_MSG)`,
	TkVariableReqbodyProcessor:               `(REQBODY_PROCESSOR)`,
	TkVariableRequestBasename:                `(REQUEST_BASENAME)`,
	TkVariableRequestBody:                    `(REQUEST_BODY)`,
	TkVariableRequestBodyLength:              `(REQUEST_BODY_LENGTH)`,
	TkVariableRequestFileName:                `(REQUEST_FILENAME)`,
	TkVariableRequestHeadersNames:            `(REQUEST_HEADERS_NAMES)`,
	TkVariableRequestLine:                    `(REQUEST_LINE)`,
	TkVariableRequestMethod:                  `(REQUEST_METHOD)`,
	TkVariableRequestProtocol:                `(REQUEST_PROTOCOL)`,
	TkVariableRequestUri:                     `(REQUEST_URI)`,
	TkVariableRequestUriRaw:                  `(REQUEST_URI_RAW)`,
	TkVariableResource:                       `(RESOURCE)`,
	TkVariableResponseBody:                   `(RESPONSE_BODY)`,
	TkVariableResponseContentLength:          `(RESPONSE_CONTENT_LENGTH)`,
	TkVariableResponseContentType:            `(RESPONSE_CONTENT_TYPE)`,
	TkVariableResponseHeadersNames:           `(RESPONSE_HEADERS_NAMES)`,
	TkVariableResponseProtocol:               `(RESPONSE_PROTOCOL)`,
	TkVariableResponseStatus:                 `(RESPONSE_STATUS)`,
	TkVariableServerAddr:                     `(SERVER_ADDR)`,
	TkVariableServerName:                     `(SERVER_NAME)`,
	TkVariableServerPort:                     `(SERVER_PORT)`,
	TkVariableSessionId:                      `(SESSIONID)`,
	TkVariableUniqueId:                       `(UNIQUE_ID)`,
	TkVariableUrlEncodedError:                `(URLENCODED_ERROR)`,
	TkVariableUserId:                         `(USERID)`,
	TkVariableWebserverErrorLog:              `(WEBSERVER_ERROR_LOG)`,
	TkVariableArgs:                           `(ARGS)`,
	TkVariableArgsPost:                       `(ARGS_POST)`,
	TkVariableArgsGet:                        `(ARGS_GET)`,
	TkVariableFilesSizes:                     `(FILES_SIZES)`,
	TkVariableFilesNames:                     `(FILES_NAMES)`,
	TkVariableFilesTmpContent:                `(FILES_TMP_CONTENT)`,
	TkVariableMatchedVarsNames:               `(MATCHED_VARS_NAMES)`,
	TkVariableMatchedVars:                    `(MATCHED_VARS)`,
	TkVariableFiles:                          `(FILES)`,
	TkVariableRequestCookies:                 `(REQUEST_COOKIES)`,
	TkVariableRequestHeaders:                 `(REQUEST_HEADERS)`,
	TkVariableResponseHeaders:                `(RESPONSE_HEADERS)`,
	TkVariableGeo:                            `(GEO)`,
	TkVariableRequestCookiesNames:            `(REQUEST_COOKIES_NAMES)`,
	TkVariableRule:                           `(RULE)`,
	TkVariableSession:                        `((SESSION))`,
	TkVariableIp:                             `((IP))`,
	TkVariableUser:                           `((USER))`,
	TkVariableStatus:                         `((STATUS))`,
	TkVariableStatusLine:                     `((STATUS_LINE))`,
	TkVariableTx:                             `(TX)`,
	TkVariableWebAppId:                       `(WEBAPPID)`,
	TkRunTimeVarBld:                          `(MODSEC_BUILD)`,
	TkRunTimeVarDur:                          `(DURATION)`,
	TkRunTimeVarEnv:                          `(ENV)`,
	TkRunTimeVarHsv:                          `(HIGHEST_SEVERITY)`,
	TkRunTimeVarRemoteUser:                   `(REMOTE_USER)`,
	TkRunTimeVarTime:                         `(TIME)`,
	TkRunTimeVarTimeDay:                      `(TIME_DAY)`,
	TkRunTimeVarTimeEpoch:                    `(TIME_EPOCH)`,
	TkRunTimeVarTimeHour:                     `(TIME_HOUR)`,
	TkRunTimeVarTimeMin:                      `(TIME_MIN)`,
	TkRunTimeVarTimeMon:                      `(TIME_MON)`,
	TkRunTimeVarTimeSec:                      `(TIME_SEC)`,
	TkRunTimeVarTimeWday:                     `(TIME_WDAY)`,
	TkRunTimeVarTimeYear:                     `(TIME_YEAR)`,
	TkRunTimeVarXml:                          `(XML)`,
	TkVarExclusion:                           `!`,
	TkVarCount:                               `&`,
	TkOperatorBeginsWith:                     `(@beginsWith)`,
	TkOperatorContains:                       `(@contains)`,
	TkOperatorContainsWord:                   `(@containsWord)`,
	TkOperatorDetectSqli:                     `(@detectSQLi)`,
	TkOperatorDetectXss:                      `(@detectXSS)`,
	TkOperatorEndsWith:                       `(@endsWith)`,
	TkOperatorEq:                             `(@eq)`,
	TkOperatorFuzzyHash:                      `(@fuzzyHash)`,
	TkOperatorGe:                             `(@ge)`,
	TkOperatorGeolookup:                      `(@geoLookup)`,
	TkOperatorGsbLookup:                      `(@gsbLookup)`,
	TkOperatorGt:                             `(@gt)`,
	TkOperatorInspectFile:                    `(@inspectFile)`,
	TkOperatorIpMatchFromFile:                `((@ipMatchF|@ipMatchFromFile))`,
	TkOperatorIpMatch:                        `(@ipMatch)`,
	TkOperatorLe:                             `(@le)`,
	TkOperatorLt:                             `(@lt)`,
	TkOperatorPmFromFile:                     `((@pmf|@pmFromFile))`,
	TkOperatorPm:                             `(@pm)`,
	TkOperatorRbl:                            `(@rbl)`,
	TkOperatorRsub:                           `(@rsub)`,
	TkOperatorRx:                             `(@rx)`,
	TkOperatorStrEq:                          `(@streq)`,
	TkOperatorStrMatch:                       `(@strmatch)`,
	TkOperatorUnconditionalMatch:             `(@unconditionalMatch)`,
	TkOperatorValidateByteRange:              `(@validateByteRange)`,
	TkOperatorValidateDtd:                    `(@validateDTD)`,
	TkOperatorValidateHash:                   `(@validateHash)`,
	TkOperatorValidateSchema:                 `(@validateSchema)`,
	TkOperatorValidateUrlEncoding:            `(@validateUrlEncoding)`,
	TkOperatorValidateUtf8Encoding:           `(@validateUtf8Encoding)`,
	TkOperatorVerifyCc:                       `(@verifyCC)`,
	TkOperatorVerifyCpf:                      `(@verifyCPF)`,
	TkOperatorVerifySsn:                      `(@verifySSN)`,
	TkOperatorVerifySvnr:                     `(@verifySVNR)`,
	TkOperatorWithin:                         `(@within)`,
	TkAuditParts:                             `[ABCDEFGHJKIZ]+`,
	TkColFreeTextSpaceComma:                  `([^,"])+`,
	TkColName:                                `[A-Za-z]+`,
	TkConfigComponentSig:                     `(SecComponentSignature)`,
	TkConfigSecServerSig:                     `(SecServerSignature)`,
	TkConfigSecWebAppId:                      `(SecWebAppId)`,
	TkConfigSecCacheTransformations:          `(SecCacheTransformations)`,
	TkConfigSecChrootDir:                     `(SecChrootDir)`,
	TkConfigConnEngine:                       `(SecConnEngine)`,
	TkConfigSecHashEngine:                    `(SecHashEngine)`,
	TkConfigSecHashKey:                       `(SecHashKey)`,
	TkConfigSecHashParam:                     `(SecHashParam)`,
	TkConfigSecHashMethodRx:                  `(SecHashMethodRx)`,
	TkConfigSecHashMethodPm:                  `(SecHashMethodPm)`,
	TkConfigContentInjection:                 `(SecContentInjection)`,
	TkConfigSecArgumentSeparator:             `(SecArgumentSeparator)`,
	TkConfigDirAuditDir:                      `(SecAuditLogStorageDir)`,
	TkConfigDirAuditDirMod:                   `(SecAuditLogDirMode)`,
	TkConfigDirAuditEng:                      `(SecAuditEngine)`,
	TkConfigDirAuditFleMod:                   `(SecAuditLogFileMode)`,
	TkConfigDirAuditLog2:                     `(SecAuditLog2)`,
	TkConfigDirAuditLog:                      `(SecAuditLog)`,
	TkConfigDirAuditLogFmt:                   `(SecAuditLogFormat)`,
	TkConfigDirAuditLogP:                     `(SecAuditLogParts)`,
	TkConfigDirAuditSts:                      `(SecAuditLogRelevantStatus)`,
	TkConfigDirAuditTpe:                      `(SecAuditLogType)`,
	TkConfigDirDebugLog:                      `(SecDebugLog)`,
	TkConfigDirDebugLvl:                      `(SecDebugLogLevel)`,
	TkConfigDirGeoDb:                         `(SecGeoLookupDb)`,
	TkConfigDirGsbDb:                         `(SecGsbLookupDb)`,
	TkConfigSecGuardianLog:                   `(SecGuardianLog)`,
	TkConfigSecInterceptOnError:              `(SecInterceptOnError)`,
	TkConfigSecConnRStateLimit:               `(SecConnReadStateLimit)`,
	TkConfigSecConnWStateLimit:               `(SecConnWriteStateLimit)`,
	TkConfigSecSensorId:                      `(SecSensorId)`,
	TkConfigSecRuleInheritance:               `(SecRuleInheritance)`,
	TkConfigSecRulePerfTime:                  `(SecRulePerfTime)`,
	TkConfigSecStreamInBodyInspection:        `(SecStreamInBodyInspection)`,
	TkConfigSecStreamOutBodyInspection:       `(SecStreamOutBodyInspection)`,
	TkConfigDirPcreMatchLimit:                `(SecPcreMatchLimit)`,
	TkConfigDirPcreMatchLimitRecursion:       `(SecPcreMatchLimitRecursion)`,
	TkConfigDirReqBody:                       `(SecRequestBodyAccess)`,
	TkConfigDirReqBodyInMemoryLimit:          `(SecRequestBodyInMemoryLimit)`,
	TkConfigDirReqBodyLimit:                  `(SecRequestBodyLimit)`,
	TkConfigDirReqBodyLimitAction:            `(SecRequestBodyLimitAction)`,
	TkConfigDirReqBodyNoFilesLimit:           `(SecRequestBodyNoFilesLimit)`,
	TkConfigDirResBody:                       `(SecResponseBodyAccess)`,
	TkConfigDirResBodyLimit:                  `(SecResponseBodyLimit)`,
	TkConfigDirResBodyLimitAction:            `(SecResponseBodyLimitAction)`,
	TkConfigDirRuleEng:                       `(SecRuleEngine)`,
	TkConfigDirSecAction:                     `(SecAction)`,
	TkConfigDirSecDefaultAction:              `(SecDefaultAction)`,
	TkConfigSecDisableBackendCompress:        `(SecDisableBackendCompression)`,
	TkConfigDirSecMarker:                     `(SecMarker)`,
	TkConfigDirUnicodeMapFile:                `(SecUnicodeMapFile)`,
	TkConfigInclude:                          `(Include)`,
	TkConfigSecCollectionTimeout:             `(SecCollectionTimeout)`,
	TkConfigSecHttpBlkey:                     `(SecHttpBlKey)`,
	TkConfigSecRemoteRules:                   `(SecRemoteRules)`,
	TkConfigSecRemoteRulesFailAction:         `(SecRemoteRulesFailAction)`,
	TkConfigSecRemoveRulesById:               `(SecRuleRemoveById)`,
	TkConfigSecRemoveRulesByMsg:              `(SecRuleRemoveByMsg)`,
	TkConfigSecRemoveRulesByTag:              `(SecRuleRemoveByTag)`,
	TkConfigSecUpdateTargetByTag:             `(SecRuleUpdateTargetByTag)`,
	TkConfigSecUpdateTargetByMsg:             `(SecRuleUpdateTargetByMsg)`,
	TkConfigSecUpdateTargetById:              `(SecRuleUpdateTargetById)`,
	TkConfigSecUpdateActionById:              `(SecRuleUpdateActionById)`,
	TkConfigUpdloadKeepFiles:                 `(SecUploadKeepFiles)`,
	TkConfigUpdloadSaveTmpFiles:              `(SecTmpSaveUploadedFiles)`,
	TkConfigUploadDir:                        `(SecUploadDir)`,
	TkConfigUploadFileLimit:                  `(SecUploadFileLimit)`,
	TkConfigUploadFileMode:                   `(SecUploadFileMode)`,
	TkConfigValueAbort:                       `(Abort)`,
	TkConfigValueDetc:                        `(DetectionOnly)`,
	TkConfigValueHttps:                       `(https)`,
	TkConfigValueNumber:                      `[0-9]+`,
	TkConfigValueOff:                         `(Off)`,
	TkConfigValueOn:                          `(On)`,
	TkConfigValueParallel:                    `(Parallel|Concurrent)`,
	TkConfigValuePath:                        `[0-9A-Za-z_\/\.\-\*\:]+`,
	TkConfigValueProcessPartial:              `(ProcessPartial)`,
	TkConfigValueReject:                      `(Reject)`,
	TkConfigValueRelevantOnly:                `(RelevantOnly)`,
	TkConfigValueSerial:                      `(Serial)`,
	TkConfigValueWarn:                        `(Warn)`,
	TkConfigXmlExternalEntity:                `(SecXmlExternalEntity)`,
	TkConfigDirResponseBodyMp:                `(SecResponseBodyMimeType)`,
	TkConfigDirResponseBodyMpClear:           `(SecResponseBodyMimeTypesClear)`,
	TkConfigDirSecArgSep:                     `(SecArgumentSeparator)`,
	TkConfigDirSecCookieFormat:               `(SecCookieFormat)`,
	TkConfigSecCookiev0Separator:             `(SecCookieV0Separator)`,
	TkConfigDirSecDataDir:                    `(SecDataDir)`,
	TkConfigDirSecStatusEngine:               `(SecStatusEngine)`,
	TkConfigDirSecTmpDir:                     `(SecTmpDir)`,
	TkDictElement:                            `([^\"|,\n \t}=]|([^\\]\\\"))+`,
	TkDictElementWithPipe:                    `[^ =\t"]+`,
	TkDictElementNoPipe:                      `[^ =\|\t"]+`,
	TkDictElementNoMacro:                     `([^\"|,%{\n \t}=]|([^\\]\\\"))+`,
	TkDictElementTwo:                         `[^\"\=, \t\r\n\\]*`,
	TkDictElementTwoQuoted:                   `[^\"\'\=\r\n\\]*`,
	TkDictElementTwo2:                        `[A-Za-z_ -\%\{\.\}\-\/]+`,
	TkDirective:                              `(SecRule)`,
	TkDirectiveSecrulescript:                 `(SecRuleScript)`,
	TkQuoteButScaped:                         `(')`,
	TkDoubleQuoteButScaped:                   `(")`,
	TkCommaButScaped:                         `(,)`,
	TkStartMacroVariable:                     `(\%\{)`,
	TkNot:                                    `!`,
	TkRemoveRuleBy:                           `[0-9A-Za-z_\/\.\-\*\:\;\]\[\$]+`,
	TkJson:                                   `(JSON)`,
	TkNative:                                 `(NATIVE)`,
	TkNewLine:                                `[\n\r]+`,
	TkEquals:                                 `(=)`,
	TkEqualsPlus:                             `(=\+)`,
	TkEqualsMinus:                            `(=\-)`,
	TkFreeText:                               freetextRegex,
	TkCharacters:                             charactersRegex,
	TkArgument:                               fmt.Sprintf(`("%s"|%s)`, freetextRegex, charactersRegex),
}

func TkRegex(i int) string {
	return TkRegexMap[i]
}

func TokenMaker(i int) func(*Scanner, *machines.Match) (interface{}, error) {
	return func(scan *Scanner, match *machines.Match) (interface{}, error) {
		return scan.Token(i, match, string(match.Bytes)), nil
	}
}

func skip(scan *Scanner, match *machines.Match) (interface{}, error) {
	return nil, nil
}

func Skip(l *Lexer, state int, reg string) {
	l.AddString(state, reg, skip)
}

// not quoted
func DirectiveToken(l *Lexer, token int, argState []int, args ...int) {
	var argsString []string
	for _, arg := range args {
		argsString = append(argsString, TkRegex(arg))
	}
	Directive(l, token, TkRegex(token), argState, argsString...)
}

// quoted
func DirectiveTokenQuote(l *Lexer, token int, argState []int, args ...int) {
	var argsString []string
	for _, arg := range args {
		argsString = append(argsString, quotedOrNot(TkRegex(arg)))
	}
	Directive(l, token, TkRegex(token), argState, argsString...)
}

func Directive(l *Lexer, tokenID int, token string, argState []int, args ...string) {
	var items []string
	sep := `[ \t]+`
	tokenString := toCaseInsensitiveRegex(token)
	items = append(items, tokenString)
	for _, arg := range args {
		items = append(items, arg)
	}
	reg := strings.Join(items, sep)

	for idx, item := range items {
		if idx == 0 {
			items[idx] = namedRegex("directive", item)
		} else {
			items[idx] = namedRegex("arg", item)
		}
	}
	reS := fmt.Sprintf("^%s$", strings.Join(items, sep))
	re := regexp.MustCompile(reS)
	names := re.SubexpNames()
	f := func(scan *Scanner, match *machines.Match) (interface{}, error) {
		var matchArgs []interface{}
		sub := re.FindSubmatch(match.Bytes)
		if len(sub) == 0 {
			return nil, fmt.Errorf("item %q not match directive %s", match.Bytes, token)
		}
		for idx, name := range names {
			if name == "arg" {
				matchArgs = append(matchArgs, removeQuotes(string(sub[idx])))
			}
		}
		for i := len(argState) - 1; i >= 0; i-- {
			scan.ToState(argState[i])
		}
		return scan.Token(tokenID, match, matchArgs...), nil
	}
	l.AddString(StateInit, reg, f)
}